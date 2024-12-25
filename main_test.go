package main

import (
	"context"
	"testing"

	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/query"
)

// MockQueryAPI simula o comportamento de QueryAPI
type MockQueryAPI struct {
	QueryFn func(ctx context.Context, query string) (api.QueryTableResult, error)
}

func (m *MockQueryAPI) Query(ctx context.Context, query string) (api.QueryTableResult, error) {
	return m.QueryFn(ctx, query)
}

// MockQueryTableResult simula o comportamento de QueryTableResult
type MockQueryTableResult struct {
	records []*query.FluxRecord // Simula os registros
	index   int                 // Índice atual
	err     error               // Erro simulado
}

// MockRecord simula o comportamento de query.FluxRecord
type MockRecord struct {
	Values map[string]interface{}
}

// Implementação do método Values
func (m *MockRecord) Values() map[string]interface{} {
	return m.Values
}

// Next avança para o próximo registro
func (m *MockQueryTableResult) Next() bool {
	m.index++
	return m.index <= len(m.records)
}

// Record retorna o registro atual
func (m *MockQueryTableResult) Record() *query.FluxRecord {
	if m.index-1 < len(m.records) {
		return m.records[m.index-1]
	}
	return nil
}

// Err retorna qualquer erro associado ao processamento
func (m *MockQueryTableResult) Err() error {
	return m.err
}

// Garante que MockQueryTableResult implementa api.QueryTableResult
var _ api.QueryTableResult = (*MockQueryTableResult)(nil)

func TestQueryResults_Success(t *testing.T) {
	mockAPI := &MockQueryAPI{
		QueryFn: func(ctx context.Context, query string) (api.QueryTableResult, error) {
			return &MockQueryTableResult{
				records: []*query.FluxRecord{
					{Values: map[string]interface{}{"field1": "value1"}},
					{Values: map[string]interface{}{"field2": "value2"}},
				},
				index: 0,
				err:   nil,
			}, nil
		},
	}

	QueryResults(mockAPI)
}
