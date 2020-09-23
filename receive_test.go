// receive_test.go.

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
	"strings"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_ReceiveJSON(t *testing.T) {

	type TestObjectClass struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
	}

	var err error
	var httpTest SimpleTest
	var test *tester.Test

	test = tester.New(t)
	httpTest = SimpleTest{
		Parameter: SimpleTestParameter{
			RequestMethod:  TestMethod,
			RequestUrl:     TestUrl,
			RequestBody:    strings.NewReader(`{"age":12345,"name":"Decode me"}`),
			RequestHandler: nil, // Is set below.
		},
	}
	objectExpected := TestObjectClass{
		Age:  12345,
		Name: "Decode me",
	}

	// Test #1. Negative Test: Not a Pointer.
	// This HTTP Handler receives an Object and checks it.
	httpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		var handlerError error
		var handlerObject TestObjectClass
		handlerError = ReceiveJSON( // <- This HTTP Handler Function is being tested.
			r,
			handlerObject,
		)
		test.MustBeAnError(handlerError)
	}
	err = PerformSimpleHttpTest(&httpTest)
	test.MustBeNoError(err)

	// Test #2. Positive Test.
	// This HTTP Handler receives an Object and checks it.
	httpTest.Parameter.RequestHandler = func(w http.ResponseWriter, r *http.Request) {
		var handlerError error
		var handlerObject TestObjectClass
		handlerError = ReceiveJSON( // <- This HTTP Handler Function is being tested.
			r,
			&handlerObject,
		)
		test.MustBeNoError(handlerError)
		test.MustBeEqual(handlerObject, objectExpected)
	}
	err = PerformSimpleHttpTest(&httpTest)
	test.MustBeNoError(err)
}
