package service

import (
	"strings"
	"testing"

	"go-basics/day28-testing/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// Register — table-driven tests (stdlib style)
// ---------------------------------------------------------------------------

func TestUserService_Register(t *testing.T) {
	tests := []struct {
		name        string
		inputName   string
		inputEmail  string
		inputPass   string
		wantErr     bool
		errContains string
	}{
		{
			name:       "valid registration",
			inputName:  "Farukul",
			inputEmail: "farukul@example.com",
			inputPass:  "secret123",
			wantErr:    false,
		},
		{
			name:        "duplicate email",
			inputName:   "Farukul",
			inputEmail:  "farukul@example.com",
			inputPass:   "secret123",
			wantErr:     true,
			errContains: "email already registered",
		},
		{
			name:        "empty password",
			inputName:   "Farukul",
			inputEmail:  "farukul2@example.com",
			inputPass:   "",
			wantErr:     true,
			errContains: "could not hash password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockUserRepo()

			// Pre-seed the duplicate email scenario
			if tt.errContains == "email already registered" {
				repo.Create(&domain.User{Name: "Other", Email: tt.inputEmail, Password: "hashed"})
			}

			svc := NewUserService(repo)
			_, err := svc.Register(tt.inputName, tt.inputEmail, tt.inputPass)

			if tt.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tt.errContains != "" && err != nil {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("expected error containing %q, got %q", tt.errContains, err.Error())
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Register — testify style (cleaner assertions)
// ---------------------------------------------------------------------------

func TestUserService_Register_Testify(t *testing.T) {
	repo := newMockUserRepo()
	svc := NewUserService(repo)

	user, err := svc.Register("Farukul", "farukul@example.com", "secret123")

	require.NoError(t, err)
	assert.Equal(t, "Farukul", user.Name)
	assert.Equal(t, "farukul@example.com", user.Email)
	assert.NotZero(t, user.ID)
}

func TestUserService_Register_DuplicateEmail_Testify(t *testing.T) {
	repo := newMockUserRepo()
	svc := NewUserService(repo)

	_, err := svc.Register("Farukul", "farukul@example.com", "secret123")
	require.NoError(t, err)

	_, err = svc.Register("Other", "farukul@example.com", "pass456")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "email already registered")
}

// ---------------------------------------------------------------------------
// Login
// ---------------------------------------------------------------------------

func TestUserService_Login(t *testing.T) {
	repo := newMockUserRepo()
	svc := NewUserService(repo)

	// Pre-register so Login has something to find
	_, err := svc.Register("Farukul", "farukul@example.com", "secret123")
	require.NoError(t, err)

	tests := []struct {
		name    string
		email   string
		pass    string
		wantErr bool
	}{
		{
			name:    "valid credentials",
			email:   "farukul@example.com",
			pass:    "secret123",
			wantErr: false,
		},
		{
			name:    "wrong password",
			email:   "farukul@example.com",
			pass:    "wrongpass",
			wantErr: true,
		},
		{
			name:    "email not found",
			email:   "nobody@example.com",
			pass:    "secret123",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := svc.Login(tt.email, tt.pass)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "invalid email or password")
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// GetProfile
// ---------------------------------------------------------------------------

func TestUserService_GetProfile(t *testing.T) {
	repo := newMockUserRepo()
	svc := NewUserService(repo)

	_, err := svc.Register("Farukul", "farukul@example.com", "secret123")
	require.NoError(t, err)

	t.Run("existing user", func(t *testing.T) {
		profile, err := svc.GetProfile(1)
		require.NoError(t, err)
		assert.Equal(t, "farukul@example.com", profile.Email)
		assert.Equal(t, "Farukul", profile.Name)
	})

	t.Run("non-existent user", func(t *testing.T) {
		_, err := svc.GetProfile(999)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

// ---------------------------------------------------------------------------
// UpdateProfile
// ---------------------------------------------------------------------------

func TestUserService_UpdateProfile(t *testing.T) {
	repo := newMockUserRepo()
	svc := NewUserService(repo)

	_, err := svc.Register("Farukul", "farukul@example.com", "secret123")
	require.NoError(t, err)

	t.Run("update name", func(t *testing.T) {
		profile, err := svc.UpdateProfile(1, "Farukul Hasan")
		require.NoError(t, err)
		assert.Equal(t, "Farukul Hasan", profile.Name)
	})

	t.Run("update non-existent user", func(t *testing.T) {
		_, err := svc.UpdateProfile(999, "Ghost")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})
}

// ---------------------------------------------------------------------------
// GetAllUsers
// ---------------------------------------------------------------------------

func TestUserService_GetAllUsers(t *testing.T) {
	repo := newMockUserRepo()
	svc := NewUserService(repo)

	t.Run("empty list", func(t *testing.T) {
		users, err := svc.GetAllUsers()
		require.NoError(t, err)
		assert.Empty(t, users)
	})

	t.Run("returns all registered users", func(t *testing.T) {
		svc.Register("Alice", "alice@example.com", "pass1")
		svc.Register("Bob", "bob@example.com", "pass2")

		users, err := svc.GetAllUsers()
		require.NoError(t, err)
		assert.Len(t, users, 2)
	})
}

// ---------------------------------------------------------------------------
// Testing helper demo — t.Helper() makes error lines point to the caller
// ---------------------------------------------------------------------------

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUserService_HelperDemo(t *testing.T) {
	repo := newMockUserRepo()
	svc := NewUserService(repo)

	_, err := svc.Register("Demo", "demo@example.com", "pass123")
	assertNoError(t, err) // error line points here, not inside assertNoError
}
