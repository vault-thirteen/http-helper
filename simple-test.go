// simple-test.go.

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
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// HTTP Methods help in testing simple HTTP Request Handlers.

// SimpleTest Type is a Type of a simple HTTP Test.
type SimpleTest struct {
	Parameter      SimpleTestParameter
	ResultExpected SimpleTestResult
	ResultReceived SimpleTestResult
}

// SimpleTestParameter Type is a Parameter of a simple HTTP Test.
type SimpleTestParameter struct {
	RequestMethod  string
	RequestUrl     string
	RequestBody    io.Reader
	RequestHandler http.HandlerFunc
}

// SimpleTestResult Type is a Result of a simple HTTP Test.
type SimpleTestResult struct {
	ResponseStatusCode int
	ResponseBodyString string
}

// PerformSimpleHttpTest Function performs the Simulation of a simple HTTP Test
// Handler. Writes the received Results into the 'ResultReceived' Field of a
// Test Object.
func PerformSimpleHttpTest(
	test *SimpleTest,
) (err error) {

	var request *http.Request
	var response *http.Response
	var responseBody []byte
	var responseRecorder *httptest.ResponseRecorder

	// Prepare Data.
	request = httptest.NewRequest(
		test.Parameter.RequestMethod,
		test.Parameter.RequestUrl,
		test.Parameter.RequestBody,
	)
	responseRecorder = httptest.NewRecorder()

	// Make a simulated Request to a HTTP Handler.
	test.Parameter.RequestHandler(responseRecorder, request)

	// Get the Response.
	response = responseRecorder.Result()
	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = response.Body.Close()
	if err != nil {
		return
	}

	// Set the Result.
	test.ResultReceived = SimpleTestResult{
		ResponseBodyString: string(responseBody),
		ResponseStatusCode: response.StatusCode,
	}
	return
}
