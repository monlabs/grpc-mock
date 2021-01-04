package stub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"

	log "github.com/golang/glog"
)

type Manager struct {
	mu    sync.RWMutex
	stubs map[string]map[string][]*Stub
}

func NewManager() *Manager {
	m := &Manager{
		stubs: make(map[string]map[string][]*Stub),
	}
	return m
}

func (m *Manager) LoadStubsFromFile(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read stubs from %s failed: %v", dir, err)
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			log.Infof("reading file %s failed: %v. skipping...", file.Name(), err)
			continue
		}

		var stubs []*Stub
		err = json.Unmarshal(data, &stubs)
		if err != nil {
			log.Infof("unmarshal stubs from file %s failed: %v. skipping...", file.Name(), err)
			continue
		}
		for _, st := range stubs {
			err = st.Validate()
			if err != nil {
				return err
			}
			m.AddStub(st)
		}
	}
	return nil
}

func (m *Manager) FindStubs(service, method string, in map[string]interface{}) []*Stub {
	m.mu.RLock()
	defer m.mu.RUnlock()
	stubs := m.stubs[service][method]
	if len(in) == 0 {
		return stubs
	}

	for _, stub := range stubs {
		if stub.Match(in) {
			return []*Stub{stub}
		}
	}

	return nil
}

func (m *Manager) AddStub(stub *Stub) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	methods := m.stubs[stub.Service]
	if methods == nil {
		methods = make(map[string][]*Stub)
		m.stubs[stub.Service] = methods
	}

	methods[stub.Method] = append(methods[stub.Method], stub)
	return nil
}

func (m *Manager) DeleteStub(service, method string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if service == "" {
		return nil
	}

	if method == "" {
		delete(m.stubs, service)
		return nil
	}

	methods := m.stubs[service]
	if methods != nil {
		delete(methods, method)
	}
	return nil
}
