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

//+build test

package httphelper

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_PerformAverageHttpTest(t *testing.T) {

	const (
		TestMethod        = "GET"
		TestURL           = "http://example.org?x=123"
		TestedHeaderName  = "X-Tester"
		TestedHeaderValue = "Me"
		TestBodyString    = "Test Body"

		ResponseStatusExpected      = http.StatusAccepted
		ResponseBodyStringExpected  = "Hello"
		ResponseHeaderNameExpected  = "X-Year"
		ResponseHeaderValueExpected = "2019"
	)

	var anAverageHttpTest AverageTest
	var aTest *tester.Test
	var err error

	aTest = tester.New(t)

	headersCommon := http.Header{}
	headersCommon.Add(TestedHeaderName, TestedHeaderValue)

	responseHeadersExpected := http.Header{}
	responseHeadersExpected.Add(ResponseHeaderNameExpected, ResponseHeaderValueExpected)

	// Test.
	anAverageHttpTest = AverageTest{
		Parameter: AverageTestParameter{
			RequestMethod:  TestMethod,
			RequestUrl:     TestURL,
			RequestHeaders: headersCommon,
			RequestBody:    strings.NewReader(TestBodyString),
			RequestHandler: nil, // Is set below.
		},
		ResultExpected: AverageTestResult{
			ResponseStatusCode: ResponseStatusExpected,
			ResponseBody:       []byte(ResponseBodyStringExpected),
			ResponseHeaders:    responseHeadersExpected,
		},
	}
	anAverageHttpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {

		var herr error
		var requestBody []byte
		var requestURLFull string

		// Verify the Request Method.
		aTest.MustBeEqual(r.Method, TestMethod)

		// Verify the incoming URL.
		requestURLFull = r.URL.Scheme + "://" + r.URL.Host + "?" + r.URL.RawQuery
		aTest.MustBeEqual(requestURLFull, TestURL)

		// Verify the incoming Header.
		inHdr := r.Header.Get(TestedHeaderName)
		aTest.MustBeEqual(inHdr, TestedHeaderValue)

		// Verify the incoming Body.
		requestBody, herr = ioutil.ReadAll(r.Body)
		aTest.MustBeNoError(herr)
		herr = r.Body.Close()
		aTest.MustBeNoError(herr)
		aTest.MustBeEqual(requestBody, []byte(TestBodyString))

		// Set the Response...

		// 1. Set an output Header.
		w.Header().Add(ResponseHeaderNameExpected, ResponseHeaderValueExpected)

		// 2. Set the Status Code.
		w.WriteHeader(ResponseStatusExpected)

		// 3. Set the Reply Body.
		_, herr = w.Write([]byte(ResponseBodyStringExpected))
		if herr != nil {
			t.FailNow()
		}
	}
	err = PerformAverageHttpTest(&anAverageHttpTest)
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(anAverageHttpTest.ResultReceived, anAverageHttpTest.ResultExpected)
}
