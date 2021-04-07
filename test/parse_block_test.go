// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"github.com/88250/lute/ast"
	"os"
	"testing"
	"time"

	"github.com/88250/lute"
	"github.com/88250/lute/parse"
	"github.com/88250/lute/util"
)

func TestBlock(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.SetVditorIR(true)
	luteEngine.ParseOptions.Mark = true
	luteEngine.ParseOptions.BlockRef = true
	luteEngine.SetKramdownIAL(true)
	luteEngine.ParseOptions.SuperBlock = true
	luteEngine.SetAutoSpace(false)
	luteEngine.SetSub(true)
	luteEngine.SetSup(true)
	luteEngine.SetGitConflict(true)

	tree := parse.Block("", util.StrToBytes("1. *foo*"), luteEngine.ParseOptions)
	list := tree.Root.FirstChild
	if ast.NodeList != list.Type {
		t.Fatalf("is not a list")
	}
	li := list.FirstChild
	if ast.NodeListItem != li.Type {
		t.Fatalf("is not a list item")
	}
	paragraph := li.FirstChild
	if ast.NodeParagraph != paragraph.Type {
		t.Fatalf("is not a paragraph")
	}
	if nil != paragraph.FirstChild {
		t.Fatalf("paragraph has children")
	}

	data, err := os.ReadFile("commonmark-spec.md")
	if nil != err {
		t.Fatalf("read test data failed: %s", err)
	}

	now := time.Now().UnixNano() / int64(time.Millisecond)
	tree = parse.Block("", data, luteEngine.ParseOptions)
	elapsed := time.Now().UnixNano()/int64(time.Millisecond) - now
	blocks := 0
	ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.WalkContinue
		}

		if !n.IsBlock() {
			return ast.WalkContinue
		}

		blocks++
		return ast.WalkContinue
	})

	t.Logf("blocks [%d], ellapsed [%d]ms", blocks, elapsed)

	now = time.Now().UnixNano() / int64(time.Millisecond)
	tree = parse.Parse("", data, luteEngine.ParseOptions)
	elapsed = time.Now().UnixNano()/int64(time.Millisecond) - now
	blocks = 0
	ast.Walk(tree.Root, func(n *ast.Node, entering bool) ast.WalkStatus {
		if !entering {
			return ast.WalkContinue
		}

		if !n.IsBlock() {
			return ast.WalkContinue
		}

		blocks++
		return ast.WalkContinue
	})

	t.Logf("blocks [%d], ellapsed [%d]ms", blocks, elapsed)
}
