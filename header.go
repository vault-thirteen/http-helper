// header.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package httphelper

import (
	"errors"
	"net/http"
	"strings"
)

// Errors
const (
	ErrHTTPHeaderNameIsNotSet   = "HTTP Header Name is not set"
	ErrHTTPHeaderNameIsNotFound = "HTTP Header Name is not found"
)

// FindHTTPHeader Function tries to find the HTTP Header with Name similar to
// the specified One. According to the Section 4.2 of RFC 2616, HTTP Header
// Names are Case insensitive. On Success, returns 'nil' and the exact Header
// Name which was found.
func FindHTTPHeader(
	r *http.Request,
	headerNameAsked string,
) (string, error) {

	var headerExists bool

	if r == nil {
		return "", errors.New(ErrNullPointer)
	}
	if len(headerNameAsked) == 0 {
		return "", errors.New(ErrHTTPHeaderNameIsNotSet)
	}

	// 1. Try the easy Way.
	_, headerExists = r.Header[headerNameAsked]
	if headerExists {
		return headerNameAsked, nil
	}

	// 2. Try the difficult Way.
	headerNameAskedLC := strings.ToLower(headerNameAsked)
	for headerName := range r.Header {
		headerNameLC := strings.ToLower(headerName)
		if headerNameLC == headerNameAskedLC {
			return headerName, nil
		}
	}

	return "", errors.New(ErrHTTPHeaderNameIsNotFound)
}

// DeleteHTTPHeader Function tries to delete a Header from the HTTP Request.
func DeleteHTTPHeader(
	r *http.Request,
	headerNameToDelete string,
) error {

	var err error
	var headerNameExact string

	// Find the Header.
	headerNameExact, err = FindHTTPHeader(r, headerNameToDelete)
	if err != nil {
		return err
	}

	// Delete the Header.
	delete(r.Header, headerNameExact)
	return nil
}
