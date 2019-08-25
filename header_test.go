////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
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

// +build test

package httphelper

import (
	"net/http"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_FindHTTPHeader(t *testing.T) {

	var aTest *tester.Test
	var err error
	var headerName string
	var request *http.Request

	aTest = tester.New(t)

	// Test #1. Null Request.
	request = nil
	headerName, err = FindHTTPHeader(request, "abc")
	aTest.MustBeAnError(err)

	// Test #2. Empty Header Name.
	request = &http.Request{}
	headerName, err = FindHTTPHeader(request, "")
	aTest.MustBeAnError(err)

	// Test #3. Existent Header with exact Name Match.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": []string{"Intergalactic Message"},
		},
	}
	headerName, err = FindHTTPHeader(request, "Content-Type")
	aTest.MustBeNoError(err)
	if headerName != "Content-Type" {
		t.FailNow()
	}

	// Test #4. Existent Header with similar Name.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": []string{"Intergalactic Message"},
		},
	}
	headerName, err = FindHTTPHeader(request, "content-type")
	aTest.MustBeNoError(err)
	if headerName != "Content-Type" {
		t.FailNow()
	}

	// Test #5. Non-existent Header.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": []string{"Intergalactic Message"},
		},
	}
	headerName, err = FindHTTPHeader(request, "X-FakeHeader")
	aTest.MustBeAnError(err)
}

func Test_DeleteHTTPHeader(t *testing.T) {

	var aTest *tester.Test
	var err error
	var request *http.Request

	aTest = tester.New(t)

	// Test #1. Test of Entry into the 'FindHTTPHeader' Function.
	request = nil
	err = DeleteHTTPHeader(request, "abc")
	aTest.MustBeAnError(err)

	// Test #2. Test of real Deletion.
	request = &http.Request{
		Header: map[string][]string{
			"Content-Type": []string{"Intergalactic Message"},
			"X-Service":    []string{"Intergalactic Service"},
		},
	}
	err = DeleteHTTPHeader(request, "x-service")
	aTest.MustBeNoError(err)
	if (len(request.Header) != 1) ||
		(len(request.Header["Content-Type"]) != 1) ||
		(request.Header["Content-Type"][0] != "Intergalactic Message") {
		t.FailNow()
	}
}
