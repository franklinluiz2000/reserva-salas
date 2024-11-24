package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIsAdmin(t *testing.T) {
	users = []User{
		{Username: "admin", Password: "1234", Admin: true},
		{Username: "user", Password: "abcd", Admin: false},
	}

	tests := []struct {
		username string
		expected bool
	}{
		{"admin", true},
		{"user", false},
		{"unknown", false},
	}

	for _, test := range tests {
		result := isAdmin(test.username)
		if result != test.expected {
			t.Errorf("isAdmin(%q) = %v; want %v", test.username, result, test.expected)
		}
	}
}

func TestHandleLogin(t *testing.T) {
	req := httptest.NewRequest("POST", "/login", nil)
	req.Form = map[string][]string{
		"username": {"admin"},
		"password": {"1234"},
	}

	users = []User{
		{Username: "admin", Password: "1234", Admin: true},
	}

	rr := httptest.NewRecorder()
	handleLogin(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("handleLogin() = %v; want %v", rr.Code, http.StatusSeeOther)
	}

	location := rr.Header().Get("Location")
	if location != "/menu" {
		t.Errorf("Redirected to %q; want /menu", location)
	}
}

func TestLeSalas(t *testing.T) {
	// Criação de arquivo temporário
	file, err := os.CreateTemp("", "salas-*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	file.WriteString("Sala1\nSala2\n")
	file.Close()

	// Chamada da função
	err = le_Salas(file.Name())
	if err != nil {
		t.Errorf("Erro ao carregar salas: %v", err)
	}

	// Validação
	if len(salas) != 2 || salas[0].Name != "Sala1" || salas[1].Name != "Sala2" {
		t.Errorf("Dados carregados incorretamente: %+v", salas)
	}
}

func TestHandleReserva(t *testing.T) {
	// Mock de dados
	salas = []Sala{{Name: "Sala1"}}
	req := httptest.NewRequest("POST", "/reserva", nil)
	req.AddCookie(&http.Cookie{Name: "username", Value: "user1"})
	req.Form = map[string][]string{
		"sala": {"Sala1"},
		"dia":  {"2024-11-24"},
		"hora": {"10:00"},
	}

	rr := httptest.NewRecorder()
	handleReserva(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("handleReserva() = %v; want %v", rr.Code, http.StatusSeeOther)
	}

	if len(salas[0].Reservas) != 1 || salas[0].Reservas[0].User != "user1" {
		t.Errorf("Reserva não foi salva corretamente: %+v", salas[0].Reservas)
	}
}

func TestLogged(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	if Logged(rr, req) {
		t.Error("Logged() retornou true para requisição sem cookie")
	}

	req.AddCookie(&http.Cookie{Name: "username", Value: ""})
	if Logged(rr, req) {
		t.Error("Logged() retornou true para cookie vazio")
	}

	// req.AddCookie(&http.Cookie{Name: "username", Value: "admin"})
	// if !Logged(rr, req) {
	// 	t.Error("Logged() retornou false para cookie válido")
	// }
}

func TestHandleMenu(t *testing.T) {
	req := httptest.NewRequest("GET", "/menu", nil)
	rr := httptest.NewRecorder()

	// Sem cookie
	handleMenu(rr, req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("handleMenu() = %v; want %v", rr.Code, http.StatusSeeOther)
	}

	// Usuário comum
	// req.AddCookie(&http.Cookie{Name: "username", Value: "user"})
	// handleMenu(rr, req)
	// if rr.Code != http.StatusOK {
	// 	t.Errorf("handleMenu() = %v; want %v", rr.Code, http.StatusOK)
	// }
}
