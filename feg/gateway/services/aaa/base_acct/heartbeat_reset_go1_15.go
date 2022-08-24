//
// Copyright 2021 The Magma Authors.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//go:build go1.15
// +build go1.15

// Package base_acct provides a client API for interacting with the
// base_acct cloud service
package base_acct

func resetHeartbeat() {
	heartbeatTicker.Reset(HeartbeatDuration)
}
