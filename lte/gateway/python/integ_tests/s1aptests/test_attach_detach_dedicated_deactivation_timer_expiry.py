"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

import ipaddress
import time
import unittest

import s1ap_types
from integ_tests.s1aptests import s1ap_wrapper
from integ_tests.s1aptests.s1ap_utils import SpgwUtil


class TestAttachDetachDedicatedDeactTmrExp(unittest.TestCase):
    """Dedicated bearer deactivation timer expiry test
    with a single UE
    """

    def setUp(self):
        """Initialize"""
        self._s1ap_wrapper = s1ap_wrapper.TestWrapper()
        self._spgw_util = SpgwUtil()

    def tearDown(self):
        """Cleanup"""
        self._s1ap_wrapper.cleanup()

    def test_attach_detach_dedicated_deactivation_timer_expiry(self):
        """attach/detach + dedicated bearer deactivation timer expiry test
        with a single UE
        """
        num_ues = 1
        self._s1ap_wrapper.configUEDevice(num_ues)

        for _ in range(num_ues):
            req = self._s1ap_wrapper.ue_req
            print(
                "********************** Running End to End attach for ",
                "UE id ",
                req.ue_id,
            )
            # Now actually complete the attach
            attach = self._s1ap_wrapper._s1_util.attach(
                req.ue_id,
                s1ap_types.tfwCmd.UE_END_TO_END_ATTACH_REQUEST,
                s1ap_types.tfwCmd.UE_ATTACH_ACCEPT_IND,
                s1ap_types.ueAttachAccept_t,
            )

            addr = attach.esmInfo.pAddr.addrInfo
            default_ip = ipaddress.ip_address(bytes(addr[:4]))

            # Wait on EMM Information from MME
            self._s1ap_wrapper._s1_util.receive_emm_info()

            print("Sleeping for 5 seconds")
            time.sleep(5)
            print(
                "********************** Adding dedicated bearer to IMSI",
                "".join([str(i) for i in req.imsi]),
            )

            # Create default flow list
            flow_list = self._spgw_util.create_default_ipv4_flows()
            self._spgw_util.create_bearer(
                "IMSI" + "".join([str(i) for i in req.imsi]),
                attach.esmInfo.epsBearerId,
                flow_list,
            )

            response = self._s1ap_wrapper.s1_util.get_response()
            self.assertEqual(
                response.msg_type, s1ap_types.tfwCmd.UE_ACT_DED_BER_REQ.value,
            )
            act_ded_ber_ctxt_req = response.cast(
                s1ap_types.UeActDedBearCtxtReq_t,
            )
            self._s1ap_wrapper.sendActDedicatedBearerAccept(
                req.ue_id, act_ded_ber_ctxt_req.bearerId,
            )

            print("Sleeping for 5 seconds")
            time.sleep(5)

            dl_flow_rules = {
                default_ip: [flow_list],
            }
            # 1 UL flow for default bearer + 1 for dedicated bearer
            num_ul_flows = 2
            # Verify if flow rules are created
            self._s1ap_wrapper.s1_util.verify_flow_rules(
                num_ul_flows, dl_flow_rules,
            )

            print(
                "********************** Deleting dedicated bearer for IMSI",
                "".join([str(i) for i in req.imsi]),
            )
            self._spgw_util.delete_bearer(
                "IMSI" + "".join([str(i) for i in req.imsi]),
                attach.esmInfo.epsBearerId,
                act_ded_ber_ctxt_req.bearerId,
            )

            # Do not send deactivate eps bearer context accept
            # TODO:Receive retransmissions once support is added
            # at s1ap tester
            response = self._s1ap_wrapper.s1_util.get_response()
            self.assertEqual(
                response.msg_type,
                s1ap_types.tfwCmd.UE_DEACTIVATE_BER_REQ.value,
            )
            deact_ded_ber_ctxt_req = response.cast(
                s1ap_types.UeDeActvBearCtxtReq_t,
            )
            print(
                "******************* Received UE_DEACTIVATE_BER_REQ with ebi ",
                deact_ded_ber_ctxt_req.bearerId,
            )

            print("Waiting for retransmissions.Sleeping for 45 seconds")
            time.sleep(45)

            print(
                "********************** Running UE detach for UE id ",
                req.ue_id,
            )
            # Now detach the UE
            self._s1ap_wrapper.s1_util.detach(
                req.ue_id,
                s1ap_types.ueDetachType_t.UE_SWITCHOFF_DETACH.value,
                False,
            )


if __name__ == "__main__":
    unittest.main()
