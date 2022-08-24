/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import * as React from 'react';
import Typography from '@mui/material/Typography';
import WifiTethering from '@mui/icons-material/WifiTethering';
import {colors} from '../theme/default';
import {makeStyles} from '@mui/styles';

const useStyles = makeStyles(() => ({
  noAccess: {
    color: colors.primary.brightGray,
    top: '50%',
    width: '520px',
    position: 'relative',
    margin: 'auto',
    textAlign: 'center',
  },
  icon: {
    width: '60px',
    height: '60px',
  },
}));

export default function ({children}: {children: React.ReactNode}) {
  const classes = useStyles();
  return (
    <Typography variant="h6" className={classes.noAccess}>
      <div>
        <WifiTethering className={classes.icon} />
      </div>
      {children}
    </Typography>
  );
}
