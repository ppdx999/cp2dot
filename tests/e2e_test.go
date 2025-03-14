package tests

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// cp2dot のバイナリをビルドし、テストケースを走査して E2E テストを実行
func TestCp2DotE2E(t *testing.T) {
	// Go バイナリのビルド
	binPath := "../cp2dot_test_bin"
	cmd := exec.Command("go", "build", "-o", binPath, "../main.go")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove(binPath) // テスト後に削除

	// `test_cases` 内のサブディレクトリを走査
	testCasesDir := "../test_cases"
	err := filepath.Walk(testCasesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ディレクトリ直下のみを処理（input.txt & output.txt のあるディレクトリ）
		if info.IsDir() && path != testCasesDir {
			inputFile := filepath.Join(path, "input.txt")
			expectedOutputFile := filepath.Join(path, "output.txt")
			argsFile := filepath.Join(path, "args.txt")

			// `input.txt`と`output.txt`の存在確認
			if _, err := os.Stat(inputFile); os.IsNotExist(err) {
				t.Errorf("Missing input file: %s", inputFile)
				return nil
			}
			if _, err := os.Stat(expectedOutputFile); os.IsNotExist(err) {
				t.Errorf("Missing output file: %s", expectedOutputFile)
				return nil
			}

			// `args.txt`の読み込み
			args := []string{}
			if _, err := os.Stat(argsFile); !os.IsNotExist(err) {
				content, _ := os.ReadFile(argsFile)
				args = strings.Fields(string(content))
			}

			// コマンド実行
			cmd := exec.Command(binPath, args...)
			inFile, _ := os.Open(inputFile)
			defer inFile.Close()
			cmd.Stdin = inFile

			var output bytes.Buffer
			cmd.Stdout = &output
			err := cmd.Run()
			if err != nil {
				t.Errorf("Execution failed for test case: %s, error: %v", path, err)
				return nil
			}

			// 期待する出力と比較
			expectedBytes, _ := os.ReadFile(expectedOutputFile)
			expected := strings.TrimSpace(string(expectedBytes))
			actual := strings.TrimSpace(output.String())

			if expected != actual {
				t.Errorf("Test case failed: %s\nExpected:\n%s\n\nGot:\n%s", path, expected, actual)
			} else {
				fmt.Printf("✅ Passed: %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		t.Fatalf("Failed to process test cases: %v", err)
	}
}
