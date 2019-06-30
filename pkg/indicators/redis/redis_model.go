/**
 * Copyright 2017 ~ 2025 the original author or authors[983708408@qq.com].
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package redis

import (
	"umc-agent/pkg/indicators"
)

// Deprecated: Metric packaging has been changed to uniform aggregation metrics,
// see: `indicators. New Metric Aggregator`.
type Redis struct {
	Meta       indicators.MetaInfo `json:"meta"`
	RedisInfos []Info              `json:"redisInfos"`
}

// Deprecated: Metric packaging has been changed to uniform aggregation metrics,
// see: `indicators. New Metric Aggregator`.
type Info struct {
	Port       string            `json:"port"`
	Properties map[string]string `json:"properties"`
}
