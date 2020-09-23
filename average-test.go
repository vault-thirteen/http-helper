// average-test.go.

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

// An average Test can emulate an HTTP Request with HTTP Method, URL, HTTP
// Headers and Body.
//
// It executes the HTTP Handler specified in the 'RequestHandler' Field,
// provides the Results of this Execution in a Filed named 'ResultReceived'.
// A User may then compare the received Results with the expected Ones.

// AverageTest Type is a Type of a average HTTP Test.
type AverageTest struct {
	Parameter      AverageTestParameter
	ResultExpected AverageTestResult
	ResultReceived AverageTestResult
}

// AverageTestParameter Type is a Parameter of a average HTTP Test.
type AverageTestParameter struct {
	RequestMethod  string
	RequestUrl     string
	RequestHeaders http.Header
	RequestBody    io.Reader
	RequestHandler http.HandlerFunc
}

// AverageTestResult Type is a Result of a average HTTP Test.
type AverageTestResult struct {
	ResponseStatusCode int
	ResponseHeaders    http.Header
	ResponseBody       []byte
}

// PerformAverageHttpTest Function performs the Simulation of an average HTTP Test
// Handler. Writes the received Results into the 'ResultReceived' Field of a
// Test Object.
func PerformAverageHttpTest(
	test *AverageTest,
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
	request.Header = test.Parameter.RequestHeaders
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
	test.ResultReceived = AverageTestResult{
		ResponseStatusCode: response.StatusCode,
		ResponseHeaders:    response.Header,
		ResponseBody:       responseBody,
	}
	return
}
