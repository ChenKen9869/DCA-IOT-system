# Details

Date : 2023-09-23 15:59:40

Directory g:\\code\\DCA-IOT-system

Total : 113 files,  14379 codes, 720 comments, 889 blanks, all 15988 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [.idea/go-backend.iml](/.idea/go-backend.iml) | XML | 9 | 0 | 0 | 9 |
| [.idea/modules.xml](/.idea/modules.xml) | XML | 8 | 0 | 0 | 8 |
| [.idea/vcs.xml](/.idea/vcs.xml) | XML | 6 | 0 | 0 | 6 |
| [README.assets/iotdb-16937282428023.svg](/README.assets/iotdb-16937282428023.svg) | XML | 1 | 0 | 0 | 1 |
| [README.assets/iotdb-16937282657945.svg](/README.assets/iotdb-16937282657945.svg) | XML | 1 | 0 | 0 | 1 |
| [README.assets/iotdb-16937282677977.svg](/README.assets/iotdb-16937282677977.svg) | XML | 1 | 0 | 0 | 1 |
| [README.assets/iotdb-16937282691659.svg](/README.assets/iotdb-16937282691659.svg) | XML | 1 | 0 | 0 | 1 |
| [README.assets/iotdb.svg](/README.assets/iotdb.svg) | XML | 1 | 0 | 0 | 1 |
| [README.assets/license-Apache 2-4EB1BA.svg](/README.assets/license-Apache%202-4EB1BA.svg) | XML | 1 | 0 | 0 | 1 |
| [README.md](/README.md) | Markdown | 283 | 0 | 139 | 422 |
| [api/common/db/database.go](/api/common/db/database.go) | Go | 52 | 0 | 7 | 59 |
| [api/common/db/device_database.go](/api/common/db/device_database.go) | Go | 37 | 0 | 8 | 45 |
| [api/common/jwt.go](/api/common/jwt.go) | Go | 35 | 0 | 7 | 42 |
| [api/common/middleware/CORS_middleware.go](/api/common/middleware/CORS_middleware.go) | Go | 18 | 0 | 3 | 21 |
| [api/common/middleware/auth_middleware.go](/api/common/middleware/auth_middleware.go) | Go | 43 | 0 | 4 | 47 |
| [api/common/middleware/cache_middleware.go](/api/common/middleware/cache_middleware.go) | Go | 1 | 1 | 2 | 4 |
| [api/common/middleware/log_middleware.go](/api/common/middleware/log_middleware.go) | Go | 56 | 0 | 5 | 61 |
| [api/rule/accepter/datasource_manager.go](/api/rule/accepter/datasource_manager.go) | Go | 44 | 0 | 14 | 58 |
| [api/rule/accepter/device_db.go](/api/rule/accepter/device_db.go) | Go | 14 | 0 | 4 | 18 |
| [api/rule/accepter/example_accepter.go](/api/rule/accepter/example_accepter.go) | Go | 67 | 0 | 6 | 73 |
| [api/rule/accepter/init.go](/api/rule/accepter/init.go) | Go | 20 | 0 | 4 | 24 |
| [api/rule/actions/action_channels.go](/api/rule/actions/action_channels.go) | Go | 7 | 0 | 4 | 11 |
| [api/rule/actions/init.go](/api/rule/actions/init.go) | Go | 30 | 0 | 6 | 36 |
| [api/rule/actions/mqtt_action.go](/api/rule/actions/mqtt_action.go) | Go | 35 | 0 | 11 | 46 |
| [api/rule/actions/ws_action.go](/api/rule/actions/ws_action.go) | Go | 38 | 0 | 7 | 45 |
| [api/rule/init.go](/api/rule/init.go) | Go | 19 | 0 | 8 | 27 |
| [api/rule/rulelog/init.go](/api/rule/rulelog/init.go) | Go | 16 | 0 | 3 | 19 |
| [api/rule/rulelog/rule_log.go](/api/rule/rulelog/rule_log.go) | Go | 5 | 0 | 3 | 8 |
| [api/rule/ruleparser/init.go](/api/rule/ruleparser/init.go) | Go | 4 | 0 | 2 | 6 |
| [api/rule/ruleparser/lexer_expression.go](/api/rule/ruleparser/lexer_expression.go) | Go | 101 | 0 | 6 | 107 |
| [api/rule/ruleparser/matcher/expression_matcher.go](/api/rule/ruleparser/matcher/expression_matcher.go) | Go | 152 | 0 | 4 | 156 |
| [api/rule/ruleparser/matcher/init.go](/api/rule/ruleparser/matcher/init.go) | Go | 6 | 0 | 3 | 9 |
| [api/rule/ruleparser/matcher/pointsurface_matcher.go](/api/rule/ruleparser/matcher/pointsurface_matcher.go) | Go | 46 | 0 | 3 | 49 |
| [api/rule/ruleparser/matcher_generator.go](/api/rule/ruleparser/matcher_generator.go) | Go | 77 | 0 | 5 | 82 |
| [api/rule/ruleparser/matcher_map.go](/api/rule/ruleparser/matcher_map.go) | Go | 4 | 0 | 4 | 8 |
| [api/rule/ruleparser/parse_action.go](/api/rule/ruleparser/parse_action.go) | Go | 30 | 4 | 5 | 39 |
| [api/rule/ruleparser/parse_condition.go](/api/rule/ruleparser/parse_condition.go) | Go | 57 | 4 | 9 | 70 |
| [api/rule/ruleparser/parse_datasource.go](/api/rule/ruleparser/parse_datasource.go) | Go | 36 | 4 | 7 | 47 |
| [api/rule/ruleparser/parser.go](/api/rule/ruleparser/parser.go) | Go | 37 | 0 | 9 | 46 |
| [api/rule/ruleparser/pre_prosess.go](/api/rule/ruleparser/pre_prosess.go) | Go | 83 | 0 | 9 | 92 |
| [api/rule/ruleparser/transform_expression..go](/api/rule/ruleparser/transform_expression..go) | Go | 80 | 0 | 5 | 85 |
| [api/rule/scheduler/init.go](/api/rule/scheduler/init.go) | Go | 11 | 0 | 3 | 14 |
| [api/rule/scheduler/rulemap.go](/api/rule/scheduler/rulemap.go) | Go | 4 | 0 | 4 | 8 |
| [api/rule/scheduler/scheduled_map.go](/api/rule/scheduler/scheduled_map.go) | Go | 7 | 0 | 4 | 11 |
| [api/server/controller/biology_controller.go](/api/server/controller/biology_controller.go) | Go | 475 | 0 | 43 | 518 |
| [api/server/controller/company_controller.go](/api/server/controller/company_controller.go) | Go | 225 | 0 | 19 | 244 |
| [api/server/controller/device_controller.go](/api/server/controller/device_controller.go) | Go | 416 | 0 | 42 | 458 |
| [api/server/controller/role_controller.go](/api/server/controller/role_controller.go) | Go | 94 | 0 | 10 | 104 |
| [api/server/controller/rule_controller.go](/api/server/controller/rule_controller.go) | Go | 179 | 0 | 21 | 200 |
| [api/server/controller/user_controller.go](/api/server/controller/user_controller.go) | Go | 92 | 0 | 12 | 104 |
| [api/server/dao/biology_dao.go](/api/server/dao/biology_dao.go) | Go | 77 | 0 | 19 | 96 |
| [api/server/dao/company_dao.go](/api/server/dao/company_dao.go) | Go | 59 | 0 | 14 | 73 |
| [api/server/dao/device_dao.go](/api/server/dao/device_dao.go) | Go | 93 | 0 | 19 | 112 |
| [api/server/dao/role_dao.go](/api/server/dao/role_dao.go) | Go | 35 | 0 | 7 | 42 |
| [api/server/dao/rule_dao.go](/api/server/dao/rule_dao.go) | Go | 36 | 0 | 8 | 44 |
| [api/server/dao/user_dao.go](/api/server/dao/user_dao.go) | Go | 25 | 0 | 6 | 31 |
| [api/server/entity/biology.go](/api/server/entity/biology.go) | Go | 49 | 0 | 7 | 56 |
| [api/server/entity/company.go](/api/server/entity/company.go) | Go | 17 | 0 | 4 | 21 |
| [api/server/entity/device.go](/api/server/entity/device.go) | Go | 31 | 0 | 5 | 36 |
| [api/server/entity/role.go](/api/server/entity/role.go) | Go | 7 | 0 | 3 | 10 |
| [api/server/entity/rule.go](/api/server/entity/rule.go) | Go | 17 | 0 | 5 | 22 |
| [api/server/entity/user.go](/api/server/entity/user.go) | Go | 12 | 0 | 2 | 14 |
| [api/server/router/biology_router.go](/api/server/router/biology_router.go) | Go | 32 | 0 | 7 | 39 |
| [api/server/router/company_router.go](/api/server/router/company_router.go) | Go | 22 | 0 | 9 | 31 |
| [api/server/router/device_router.go](/api/server/router/device_router.go) | Go | 34 | 0 | 11 | 45 |
| [api/server/router/monitor_router.go](/api/server/router/monitor_router.go) | Go | 13 | 0 | 7 | 20 |
| [api/server/router/role_router.go](/api/server/router/role_router.go) | Go | 16 | 0 | 6 | 22 |
| [api/server/router/rule_router.go](/api/server/router/rule_router.go) | Go | 19 | 0 | 7 | 26 |
| [api/server/router/user_router.go](/api/server/router/user_router.go) | Go | 15 | 1 | 5 | 21 |
| [api/server/service/biology_service.go](/api/server/service/biology_service.go) | Go | 263 | 209 | 26 | 498 |
| [api/server/service/company_service.go](/api/server/service/company_service.go) | Go | 283 | 85 | 18 | 386 |
| [api/server/service/device_service.go](/api/server/service/device_service.go) | Go | 304 | 167 | 25 | 496 |
| [api/server/service/role_service.go](/api/server/service/role_service.go) | Go | 78 | 37 | 6 | 121 |
| [api/server/service/rule_service.go](/api/server/service/rule_service.go) | Go | 149 | 80 | 17 | 246 |
| [api/server/service/sensor_service.go](/api/server/service/sensor_service.go) | Go | 67 | 29 | 7 | 103 |
| [api/server/service/user_service.go](/api/server/service/user_service.go) | Go | 103 | 50 | 8 | 161 |
| [api/server/tools/server/message.go](/api/server/tools/server/message.go) | Go | 13 | 0 | 1 | 14 |
| [api/server/tools/server/request.go](/api/server/tools/server/request.go) | Go | 100 | 0 | 15 | 115 |
| [api/server/tools/server/response.go](/api/server/tools/server/response.go) | Go | 43 | 0 | 11 | 54 |
| [api/server/tools/util/hex_to_dec.go](/api/server/tools/util/hex_to_dec.go) | Go | 9 | 0 | 2 | 11 |
| [api/server/tools/util/is_float_qual.go](/api/server/tools/util/is_float_qual.go) | Go | 6 | 0 | 4 | 10 |
| [api/server/tools/util/parse_date.go](/api/server/tools/util/parse_date.go) | Go | 9 | 0 | 3 | 12 |
| [api/server/tools/util/stack.go](/api/server/tools/util/stack.go) | Go | 17 | 0 | 7 | 24 |
| [api/server/tools/util/string_translate.go](/api/server/tools/util/string_translate.go) | Go | 66 | 0 | 5 | 71 |
| [api/server/vo/biology_info.go](/api/server/vo/biology_info.go) | Go | 35 | 0 | 8 | 43 |
| [api/server/vo/biology_portabledevice.go](/api/server/vo/biology_portabledevice.go) | Go | 8 | 0 | 1 | 9 |
| [api/server/vo/company_info.go](/api/server/vo/company_info.go) | Go | 6 | 0 | 1 | 7 |
| [api/server/vo/company_tree_node.go](/api/server/vo/company_tree_node.go) | Go | 9 | 0 | 2 | 11 |
| [api/server/vo/fio_data.go](/api/server/vo/fio_data.go) | Go | 11 | 0 | 2 | 13 |
| [api/server/vo/fixed_device.go](/api/server/vo/fixed_device.go) | Go | 15 | 0 | 3 | 18 |
| [api/server/vo/new_collar.go](/api/server/vo/new_collar.go) | Go | 14 | 0 | 1 | 15 |
| [api/server/vo/own_assets.go](/api/server/vo/own_assets.go) | Go | 29 | 0 | 4 | 33 |
| [api/server/vo/portable_device.go](/api/server/vo/portable_device.go) | Go | 11 | 0 | 2 | 13 |
| [api/server/vo/position_collar.go](/api/server/vo/position_collar.go) | Go | 9 | 0 | 2 | 11 |
| [api/server/vo/rule_info.go](/api/server/vo/rule_info.go) | Go | 12 | 0 | 3 | 15 |
| [api/sys/gis/geo/geoalgorithm/spatial_analysis_2d.go](/api/sys/gis/geo/geoalgorithm/spatial_analysis_2d.go) | Go | 28 | 0 | 3 | 31 |
| [api/sys/gis/geo/geocontainer/coordinate.go](/api/sys/gis/geo/geocontainer/coordinate.go) | Go | 29 | 0 | 6 | 35 |
| [api/sys/gis/geo/geocontainer/geo_position_2d.go](/api/sys/gis/geo/geocontainer/geo_position_2d.go) | Go | 23 | 0 | 6 | 29 |
| [api/sys/iot/monitor/connection.go](/api/sys/iot/monitor/connection.go) | Go | 89 | 16 | 11 | 116 |
| [api/sys/iot/sensor/collar_message.go](/api/sys/iot/sensor/collar_message.go) | Go | 75 | 0 | 6 | 81 |
| [api/sys/iot/sensor/collar_msg_http.go](/api/sys/iot/sensor/collar_msg_http.go) | Go | 9 | 0 | 2 | 11 |
| [api/sys/iot/sensor/collections.go](/api/sys/iot/sensor/collections.go) | Go | 14 | 0 | 4 | 18 |
| [api/sys/iot/sensor/fio_message.go](/api/sys/iot/sensor/fio_message.go) | Go | 77 | 0 | 8 | 85 |
| [api/sys/iot/sensor/position_collar_message.go](/api/sys/iot/sensor/position_collar_message.go) | Go | 68 | 0 | 7 | 75 |
| [config/application.yaml](/config/application.yaml) | YAML | 43 | 0 | 0 | 43 |
| [docs/docs.go](/docs/docs.go) | Go | 2,860 | 3 | 5 | 2,868 |
| [docs/swagger.json](/docs/swagger.json) | JSON | 2,844 | 0 | 0 | 2,844 |
| [docs/swagger.yaml](/docs/swagger.yaml) | YAML | 1,911 | 0 | 1 | 1,912 |
| [go.mod](/go.mod) | Go Module File | 91 | 0 | 5 | 96 |
| [go.sum](/go.sum) | Go Checksum File | 803 | 0 | 1 | 804 |
| [main.go](/main.go) | Go | 43 | 30 | 9 | 82 |
| [scripts/build.bat](/scripts/build.bat) | Batch | 8 | 0 | 2 | 10 |
| [scripts/example_device.py](/scripts/example_device.py) | Python | 21 | 0 | 4 | 25 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)