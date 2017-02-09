// Copyright © 2016 Abcum Ltd
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

package sql

func (p *parser) parseDefineLoginStatement() (stmt *DefineLoginStatement, err error) {

	stmt = &DefineLoginStatement{}

	if stmt.User, err = p.parseName(); err != nil {
		return nil, err
	}

	if _, _, err = p.shouldBe(ON); err != nil {
		return nil, err
	}

	if stmt.Kind, _, err = p.shouldBe(NAMESPACE, DATABASE); err != nil {
		return nil, err
	}

	if p.is(stmt.Kind, NAMESPACE) {
		if stmt.KV, stmt.NS, stmt.DB, err = p.o.get(AuthNS); err != nil {
			return nil, err
		}
	}

	if p.is(stmt.Kind, DATABASE) {
		if stmt.KV, stmt.NS, stmt.DB, err = p.o.get(AuthDB); err != nil {
			return nil, err
		}
	}

	if _, _, err := p.shouldBe(PASSWORD); err != nil {
		return nil, err
	}

	if stmt.Pass, err = p.parseBcrypt(); err != nil {
		return nil, err
	}

	if stmt.Code, err = p.parseRand(); err != nil {
		return nil, err
	}

	if _, _, err = p.shouldBe(EOF, SEMICOLON); err != nil {
		return nil, err
	}

	return

}

func (p *parser) parseRemoveLoginStatement() (stmt *RemoveLoginStatement, err error) {

	stmt = &RemoveLoginStatement{}

	if stmt.User, err = p.parseName(); err != nil {
		return nil, err
	}

	if _, _, err = p.shouldBe(ON); err != nil {
		return nil, err
	}

	if stmt.Kind, _, err = p.shouldBe(NAMESPACE, DATABASE); err != nil {
		return nil, err
	}

	if p.is(stmt.Kind, NAMESPACE) {
		if stmt.KV, stmt.NS, stmt.DB, err = p.o.get(AuthNS); err != nil {
			return nil, err
		}
	}

	if p.is(stmt.Kind, DATABASE) {
		if stmt.KV, stmt.NS, stmt.DB, err = p.o.get(AuthDB); err != nil {
			return nil, err
		}
	}

	if _, _, err = p.shouldBe(EOF, SEMICOLON); err != nil {
		return nil, err
	}

	return

}