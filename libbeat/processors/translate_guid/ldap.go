// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package translate_guid

import (
	"crypto/tls"
	"fmt"
	"strings"
	"sync"

	"github.com/go-ldap/ldap/v3"
)

// ldapClient manages a single reusable LDAP connection
type ldapClient struct {
	conn *ldap.Conn
	mu   sync.Mutex
	*ldapConfig
}

type ldapConfig struct {
	address         string
	baseDN          string
	username        string
	password        string
	searchTimeLimit int
	tlsConfig       *tls.Config
}

// newLDAPClient initializes a new ldapClient with a single connection
func newLDAPClient(config *ldapConfig) (*ldapClient, error) {
	client := &ldapClient{ldapConfig: config}

	// Establish initial connection
	if err := client.connect(); err != nil {
		return nil, err
	}

	return client, nil
}

// connect establishes a new connection to the LDAP server
func (client *ldapClient) connect() error {
	client.mu.Lock()
	defer client.mu.Unlock()

	// Connect with or without TLS based on configuration
	var conn *ldap.Conn
	var err error
	if client.tlsConfig != nil {
		conn, err = ldap.DialTLS("tcp", client.address, client.tlsConfig)
	} else {
		conn, err = ldap.Dial("tcp", client.address)
	}
	if err != nil {
		return fmt.Errorf("failed to dial LDAP server: %v", err)
	}

	if client.password != "" {
		err = conn.Bind(client.username, client.password)
	} else {
		err = conn.UnauthenticatedBind(client.username)
	}

	if err != nil {
		conn.Close()
		return fmt.Errorf("failed to bind to LDAP server: %v", err)
	}

	client.conn = conn
	return nil
}

// reconnect checks the connection's health and reconnects if necessary
func (client *ldapClient) reconnect() error {
	client.mu.Lock()
	defer client.mu.Unlock()

	// Check if the connection is still alive
	if client.conn.IsClosing() {
		return client.connect()
	}
	return nil
}

// findObjectByGUID searches for an AD object by GUID and returns its Common Name (CN)
func (client *ldapClient) findObjectByGUID(objectGUID string) (string, error) {
	// Ensure the connection is alive or reconnect if necessary
	if err := client.reconnect(); err != nil {
		return "", fmt.Errorf("failed to reconnect: %v", err)
	}

	client.mu.Lock()
	defer client.mu.Unlock()

	// Format the GUID filter and perform the search
	filter := fmt.Sprintf("(objectGUID=%s)", encodeGUID(objectGUID))
	searchRequest := ldap.NewSearchRequest(
		client.baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 1, client.searchTimeLimit, false,
		filter, []string{"cn"}, nil,
	)

	// Execute search
	result, err := client.conn.Search(searchRequest)
	if err != nil {
		return "", fmt.Errorf("search failed: %v", err)
	}
	if len(result.Entries) == 0 {
		return "", fmt.Errorf("no entries found for GUID %s", objectGUID)
	}

	// Retrieve the CN attribute
	cn := result.Entries[0].GetAttributeValue("cn")
	return cn, nil
}

// encodeGUID converts a GUID into LDAP filter format
func encodeGUID(guid string) string {
	return fmt.Sprintf("\\%s", strings.Trim(guid, "{}"))
}

// close closes the LDAP connection
func (client *ldapClient) close() {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.conn != nil {
		client.conn.Close()
	}
}
