// (c) Copyright 2016 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rules

import (
	"go/ast"

	"github.com/GoASTScanner/gas"
)

type usesWeakCryptography struct {
	gas.MetaData
	blacklist map[string][]string
}

func (r *usesWeakCryptography) ID() string {
	return r.MetaData.ID
}

func (r *usesWeakCryptography) Match(n ast.Node, c *gas.Context) (*gas.Issue, error) {
	for pkg, funcs := range r.blacklist {
		if _, matched := gas.MatchCallByPackage(n, c, pkg, funcs...); matched {
			return gas.NewIssue(c, n, r.ID(), r.What, r.Severity, r.Confidence), nil
		}
	}
	return nil, nil
}

// NewUsesWeakCryptography detects uses of des.* md5.* or rc4.*
func NewUsesWeakCryptography(id string, conf gas.Config) (gas.Rule, []ast.Node) {
	calls := make(map[string][]string)
	calls["crypto/des"] = []string{"NewCipher", "NewTripleDESCipher"}
	calls["crypto/md5"] = []string{"New", "Sum"}
	calls["crypto/rc4"] = []string{"NewCipher"}
	rule := &usesWeakCryptography{
		blacklist: calls,
		MetaData: gas.MetaData{
			ID:         id,
			Severity:   gas.Medium,
			Confidence: gas.High,
			What:       "Use of weak cryptographic primitive",
		},
	}
	return rule, []ast.Node{(*ast.CallExpr)(nil)}
}
