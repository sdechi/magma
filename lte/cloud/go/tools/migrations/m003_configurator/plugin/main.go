/*
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

package main

import (
	"fmt"

	"magma/lte/cloud/go/tools/migrations/m003_configurator/plugin/types"
	"magma/orc8r/cloud/go/sqorc"
	"magma/orc8r/cloud/go/tools/migrations/m003_configurator/migration"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
)

const policyTable = "policydb"
const baseNameTable = "base_names"

func main() {}

type plugin struct{}

func GetPlugin() migration.ConfiguratorMigrationPlugin {
	return &plugin{}
}

func (*plugin) GetConfigMigrators() []migration.ConfigMigrator {
	return []migration.ConfigMigrator{
		&cellularNetworkMigrator{},
		&cellularGatewayMigrator{},
	}
}

func (*plugin) RunCustomMigrations(
	sc *squirrel.StmtCache,
	builder sqorc.StatementBuilder,
	migratedGatewayMetasByNetwork map[string]map[string]migration.MigratedGatewayMeta,
) error {
	nids := funk.Keys(migratedGatewayMetasByNetwork).([]string)

	if err := migratePolicydb(sc, builder, nids); err != nil {
		return errors.WithStack(err)
	}

	if err := migrateSubscriberdb(sc, builder, nids); err != nil {
		return errors.WithStack(err)
	}

	if err := migrateEnodebs(sc, builder, migratedGatewayMetasByNetwork); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// migrators

type cellularNetworkMigrator struct{}
type cellularGatewayMigrator struct{}

func (*cellularNetworkMigrator) GetType() string {
	return "cellular_network"
}

func (*cellularNetworkMigrator) ToNewConfig(oldConfig []byte) ([]byte, error) {
	oldMsg := &types.CellularNetworkConfig{}
	err := migration.Unmarshal(oldConfig, oldMsg)
	if err != nil {
		return nil, err
	}

	newModel := &types.NetworkCellularConfigs{}
	migration.FillIn(oldMsg, newModel)
	newModel.FegNetworkID = oldMsg.FegNetworkId
	newModel.Epc.RelayEnabled = oldMsg.Epc.RelayEnabled

	if oldMsg.Epc.NetworkServices != nil {
		for _, serviceEnum := range oldMsg.Epc.NetworkServices {
			serviceName, ok := networkServiceEnumToNameMap[serviceEnum]
			if !ok {
				return nil, fmt.Errorf("Unknown network service enum: %s", serviceEnum)
			}
			newModel.Epc.NetworkServices = append(newModel.Epc.NetworkServices, serviceName)
		}
	}

	return newModel.MarshalBinary()
}

var networkServiceEnumToNameMap = map[types.NetworkEPCConfig_NetworkServices]string{
	types.NetworkEPCConfig_METERING:    "metering",
	types.NetworkEPCConfig_DPI:         "dpi",
	types.NetworkEPCConfig_ENFORCEMENT: "policy_enforcement",
}

func (*cellularGatewayMigrator) GetType() string {
	return "cellular_gateway"
}

func (*cellularGatewayMigrator) ToNewConfig(oldConfig []byte) ([]byte, error) {
	oldMsg := &types.CellularGatewayConfig{}
	err := migration.Unmarshal(oldConfig, oldMsg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	newModel := &types.GatewayCellularConfigs{}
	migration.FillIn(oldMsg, newModel)
	return newModel.MarshalBinary()

}
