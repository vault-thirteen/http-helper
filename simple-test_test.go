// simple-test_test.go.

//+build test

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
	"net/http"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_PerformSimpleHttpTest(t *testing.T) {

	var aSimpleHttpTest SimpleTest
	var aTest *tester.Test
	var err error

	aTest = tester.New(t)

	// Test #1. Positive.
	aSimpleHttpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  "GET",
			RequestUrl:     "http://example.org",
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: SimpleTestResult{
			ResponseBodyString: "Hello",
			ResponseStatusCode: http.StatusAccepted,
		},
	}
	aSimpleHttpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		var herr error
		w.WriteHeader(http.StatusAccepted)
		_, herr = w.Write([]byte("Hello"))
		if herr != nil {
			t.FailNow()
		}
	}
	err = PerformSimpleHttpTest(&aSimpleHttpTest)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(aSimpleHttpTest.ResultReceived, aSimpleHttpTest.ResultExpected)

	// Test #2. Negative.
	aSimpleHttpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  "GET",
			RequestUrl:     "http://example.org",
			RequestBody:    nil,
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: SimpleTestResult{
			ResponseBodyString: "Hello",
			ResponseStatusCode: http.StatusAccepted,
		},
	}
	aSimpleHttpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		var herr error
		w.WriteHeader(http.StatusBadRequest)
		_, herr = w.Write([]byte("Bye-bye"))
		if herr != nil {
			t.FailNow()
		}
	}
	err = PerformSimpleHttpTest(&aSimpleHttpTest)
	aTest.MustBeNoError(err)
	aTest.MustBeDifferent(aSimpleHttpTest.ResultReceived, aSimpleHttpTest.ResultExpected)
}
