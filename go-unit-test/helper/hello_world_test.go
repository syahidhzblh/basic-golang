package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Benchmark Test
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Syahid")
	}
}

// Sub Benchmark
func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("Syahid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Syahid")
		}
	})

	b.Run("Hisbul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hisbul")
		}
	})
}

//Table Benchmark

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		Name    string
		Request string
	}{
		{
			Name:    "Syahid",
			Request: "Hisbul",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.Request)
			}
		})
	}
}

// Prefer using TableTest for Unit Test. Cause, mostly programmer golang using TableTest
func TestHelloWorldTable(t *testing.T) {
	//Slice Struct for Value on SubTest
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "Syahid",
			Request:  "Syahid",
			Expected: "Hello Syahid",
		},
		{
			Name:     "Hisbul",
			Request:  "Hisbul",
			Expected: "Hello Hisbul",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Name)
			require.Equal(t, test.Expected, result, "Result Must Be Hello %+v", test.Request)
		})
	}
}

func TestSubTest(t *testing.T) {

	t.Run("Syahid", func(t *testing.T) {
		result := HelloWorld("Syahid")
		require.Equal(t, "Hello Syahid", result, "Result Must Be 'Hello Syahid'")
	})
	t.Run("Hisbul", func(t *testing.T) {
		result := HelloWorld("Hisbul")
		require.Equal(t, "Hello Hisbul", result, "Result Must Be 'Hello Hisbul'")
	})
}

// TestMain hanya dieksekusi sekali per package
func TestMain(m *testing.M) {
	//Before Test (Bisa digunakan untuk test koneksi database,dll. Sebelum Unit test dijalankan)
	fmt.Println("Before Unit Test")

	m.Run()

	//After Unit Test
	fmt.Println("After Unit Test")
}

// t.Error same like t.Fail , code after test will be execute
func TestHelloWorldSyahid(t *testing.T) {
	result := HelloWorld("Syahid")

	if result != "Hello Syahid" {
		t.Error("Result must be 'Hello Syahid'")
	}

	fmt.Println("TestHelloWorldSyahid Done")
}

// t.Fatal same like t.FailNow. After testing fail, code after test will not be executed
func TestHelloWorldHisbul(t *testing.T) {
	result := HelloWorld("Hisbul")

	if result != "Hello Hisbul" {
		t.Fatal("Result must be 'Hello Hisbul'")
	}

	fmt.Println("TestHelloWorldHisbul Done")

}

// assert same like t.Fail , code after test will be execute
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Syahid")

	assert.Equal(t, "Hello Syahid", result, "Result must be 'Hello Syahid'")
	fmt.Println("Testing TestHelloWorldAssert Done")
}

// require same like t.FailNow. After testing fail, code after test will not be executed
func TestHelloWorldRequier(t *testing.T) {
	result := HelloWorld("Syahid")

	require.Equal(t, "Hello Syahid", result, "Result must be 'Hello Syahid'")
	fmt.Println("Testing TestHelloWorldAssert Done")
}

// Skip Test digunakan untuk kondisi tertentu, akan mengecek kondisi dulu. Jika kondisinya tidak sesuai akan melakukan Skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on linux")
	}

	result := HelloWorld("Syahid")
	require.Equal(t, "Hello Syahid", result, "Result must be 'Hello Syahid'")
}
