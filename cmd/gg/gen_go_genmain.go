// Code generated by genmain. DO NOT EDIT.

package main

const goMain = "\"package main\\n\\nimport (\\n\\t\\\"fmt\\\"\\n\\t\\\"os\\\"\\n\\t\\\"os/exec\\\"\\n\\t\\\"path/filepath\\\"\\n\\t\\\"runtime\\\"\\n\\t\\\"strings\\\"\\n)\\n\\nfunc main() {\\n\\truntime.LockOSThread()\\n\\n\\tfmt.Fprintf(os.Stderr, \\\"go %v\\\\n\\\", strings.Join(os.Args[1:], \\\" \\\"))\\n\\tsep := string(filepath.ListSeparator)\\n\\tself, err := os.Executable()\\n\\tif err != nil {\\n\\t\\tfmt.Fprintf(os.Stderr, \\\"failed to derive path to self: %v\\\", err)\\n\\t\\tos.Exit(1)\\n\\t}\\n\\tselfDir := filepath.Dir(self)\\n\\tvar path []string\\n\\tfor _, p := range strings.Split(os.Getenv(\\\"PATH\\\"), sep) {\\n\\t\\tif p != selfDir {\\n\\t\\t\\tpath = append(path, p)\\n\\t\\t}\\n\\t}\\n\\n\\tif err := os.Setenv(\\\"PATH\\\", strings.Join(path, sep)); err != nil {\\n\\t\\tfmt.Fprintf(os.Stderr, \\\"failed to set PATH: %v\\\", err)\\n\\t\\tos.Exit(1)\\n\\t}\\n\\n\\tcmd := exec.Command(\\\"go\\\", os.Args[1:]...)\\n\\tcmd.Stdout = os.Stdout\\n\\tcmd.Stderr = os.Stderr\\n\\tcmd.Stdin = os.Stdin\\n\\n\\tif err := cmd.Run(); err != nil {\\n\\t\\tos.Exit(1)\\n\\t}\\n}\\n\""
