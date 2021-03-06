/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package http_test

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/dubbogo/dubbo-go-proxy/pkg/common/extension"
	"github.com/dubbogo/dubbo-go-proxy/pkg/config"
	"github.com/dubbogo/dubbo-go-proxy/pkg/context"
	httpCtx "github.com/dubbogo/dubbo-go-proxy/pkg/context/http"
	_ "github.com/dubbogo/dubbo-go-proxy/pkg/filter"
	"github.com/dubbogo/dubbo-go-proxy/pkg/model"
	"github.com/dubbogo/dubbo-go-proxy/pkg/router"
)

func getMockAPI(verb config.HTTPVerb, urlPattern string, filters ...string) router.API {
	inbound := config.InboundRequest{}
	integration := config.IntegrationRequest{RequestType: config.DubboRequest}
	method := config.Method{
		OnAir:              true,
		HTTPVerb:           verb,
		InboundRequest:     inbound,
		IntegrationRequest: integration,
		Filters:            filters,
	}
	return router.API{
		URLPattern: urlPattern,
		Method:     method,
	}
}

func TestBuildContext(t *testing.T) {
	extension.SetFilterFunc("a", func(ctx context.Context) { ctx.Next() })
	extension.SetFilterFunc("b", func(ctx context.Context) { ctx.Next() })
	extension.SetFilterFunc("c", func(ctx context.Context) { ctx.Next() })
	ctx := httpCtx.HttpContext{
		FilterChains: []model.FilterChain{},
		BaseContext:  context.NewBaseContext(),
	}
	ctx.API(getMockAPI(config.MethodPost, "/mock/test", "a", "b", "c"))
	ctx.BuildFilters()

	assert.Equal(t, len(ctx.Filters), 4)
}
