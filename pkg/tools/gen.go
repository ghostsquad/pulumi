// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/pulumi/pulumi-fabric/pkg/util/contract"
)

// GenWriter adds some convenient helpers atop a buffered writer.
type GenWriter struct {
	tool string        // the name of the code-generator.
	f    *os.File      // the file being written to.
	buff *bytes.Buffer // the buffer (if there is no file).
	w    *bufio.Writer // the buffered writer used to emit code.
}

func NewGenWriter(tool string, file string) (*GenWriter, error) {
	// If the file is non-empty, open up a writer and overwrite whatever file contents already exist.
	if file != "" {
		f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return nil, err
		}
		return &GenWriter{tool: tool, f: f, w: bufio.NewWriter(f)}, nil
	}

	// Otherwise, we are emitting into an in-memory buffer.
	var buff bytes.Buffer
	return &GenWriter{tool: tool, buff: &buff, w: bufio.NewWriter(&buff)}, nil
}

// Flush explicitly flushes the writer's pending writes.
func (g *GenWriter) Flush() error {
	return g.w.Flush()
}

// Close flushes and closes the underlying writer.
func (g *GenWriter) Close() error {
	err := g.w.Flush()
	contract.IgnoreError(err)
	if g.f != nil {
		return g.f.Close()
	}
	return nil
}

// Writefmt wraps the bufio.Writer.WriteString function, but also performs fmt.Sprintf-style formatting.
func (g *GenWriter) Writefmt(msg string, args ...interface{}) {
	_, err := g.w.WriteString(fmt.Sprintf(msg, args...))
	contract.IgnoreError(err)
}

// Writefmtln wraps the bufio.Writer.WriteString function, performing fmt.Sprintf-style formatting and appending \n.
func (g *GenWriter) Writefmtln(msg string, args ...interface{}) {
	g.Writefmt(msg+"\n", args...)
}

// EmitHeaderWarning emits the standard "WARNING" into a generated file.
func (g *GenWriter) EmitHeaderWarning() {
	g.Writefmtln("// *** WARNING: this file was generated by %v. ***", g.tool)
	g.Writefmtln("// *** Do not edit by hand unless you're certain you know what you are doing! ***")
	g.Writefmtln("")
}

// Buffer returns whatever has been written to the in-memory buffer (in non-file cases).
func (g *GenWriter) Buffer() string {
	return g.buff.String()
}
