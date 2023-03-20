package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Fachrul",
			request: "Fachrul",
		},
		{
			name:    "Aziz",
			request: "Aziz",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
func BenchmarkSub(b *testing.B) {
	b.Run("Fachrul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fachrul")
		}
	})
	b.Run("Aziz", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Aziz")
		}
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "fachrul",
			request:  "fachrul",
			expected: "Hello fachrul",
		},
		{

			name:     "daffara",
			request:  "daffara",
			expected: "Hello daffara",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result Must be "+test.expected)
		})
	}
}
func TestSubTest(a *testing.T) {
	a.Run("desyana", func(t *testing.T) {
		result := HelloWorld("desyana")
		require.Equal(a, "Hello desyana", result, "Result must be 'Hello desyana")
	})
	a.Run("intani", func(t *testing.T) {
		result := HelloWorld("intani")
		require.Equal(a, "Hello intani", result, "Result must be 'Hello desyana'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" { // <= jika sistem MacOs bernama darwin. tetapi jika windows tetap windows
		t.Skip("tidak bisa dijalankan di MacOs")
	}
	result := HelloWorld("Daffara")
	fmt.Println(result)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Fachrul")
	require.Equal(t, "Hello Fachrul", result, "Result must be 'Hello Fachrul'")
	fmt.Println("TestHelloWorld with Require Done")
}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fachrul Aziz")
	assert.Equal(t, "Hello Fachrul Aziz", result, "Result must be 'Hello Fachrul Aziz'")
	fmt.Println("TestHelloWorld with Assert Done")
}
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fachrul Aziz")
	if result != "Hello Fachrul Aziz" {
		// error menggunakan error
		t.Error("Result must be 'Hello Fachrul Aziz'")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldNaura(t *testing.T) {
	result := HelloWorld("Naura Almeera Fachrul")
	if result != "Hello Naura Almeera Fachrul" {
		// error fatal
		t.Fatal("Result must be 'Hello Naura Almeera Fachrul'")
	}
	fmt.Println("TestHelloWorldNaura Done")
}
