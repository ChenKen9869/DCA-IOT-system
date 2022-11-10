# Details

Date : 2022-10-04 16:32:00

Directory g:\\code\\go-backend

Total : 91 files,  10029 codes, 695 comments, 589 blanks, all 11313 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [.idea/go-backend.iml](/.idea/go-backend.iml) | XML | 9 | 0 | 0 | 9 |
| [.idea/modules.xml](/.idea/modules.xml) | XML | 8 | 0 | 0 | 8 |
| [.idea/vcs.xml](/.idea/vcs.xml) | XML | 6 | 0 | 0 | 6 |
| [README.md](/README.md) | Markdown | 24 | 0 | 19 | 43 |
| [cache/dao_cache.go](/cache/dao_cache.go) | Go | 26 | 6 | 10 | 42 |
| [cache/lru_cache.go](/cache/lru_cache.go) | Go | 1 | 1 | 1 | 3 |
| [common/database.go](/common/database.go) | Go | 49 | 2 | 7 | 58 |
| [common/device_database.go](/common/device_database.go) | Go | 36 | 2 | 9 | 47 |
| [common/jwt.go](/common/jwt.go) | Go | 35 | 1 | 12 | 48 |
| [common/redis.go](/common/redis.go) | Go | 25 | 2 | 8 | 35 |
| [config/application.yaml](/config/application.yaml) | YAML | 36 | 0 | 0 | 36 |
| [controller/biology_controller.go](/controller/biology_controller.go) | Go | 370 | 23 | 41 | 434 |
| [controller/company_controller.go](/controller/company_controller.go) | Go | 196 | 7 | 12 | 215 |
| [controller/device_controller.go](/controller/device_controller.go) | Go | 258 | 13 | 34 | 305 |
| [controller/fence_controller.go](/controller/fence_controller.go) | Go | 128 | 7 | 8 | 143 |
| [controller/role_controller.go](/controller/role_controller.go) | Go | 1 | 0 | 2 | 3 |
| [controller/user_controller.go](/controller/user_controller.go) | Go | 67 | 1 | 12 | 80 |
| [dao/biology_dao.go](/dao/biology_dao.go) | Go | 69 | 5 | 19 | 93 |
| [dao/company_dao.go](/dao/company_dao.go) | Go | 51 | 0 | 12 | 63 |
| [dao/device_dao.go](/dao/device_dao.go) | Go | 73 | 8 | 18 | 99 |
| [dao/fence_dao.go](/dao/fence_dao.go) | Go | 40 | 3 | 11 | 54 |
| [dao/user_dao.go](/dao/user_dao.go) | Go | 19 | 3 | 8 | 30 |
| [docs/docs.go](/docs/docs.go) | Go | 1,981 | 3 | 5 | 1,989 |
| [docs/swagger.json](/docs/swagger.json) | JSON | 1,965 | 0 | 0 | 1,965 |
| [docs/swagger.yaml](/docs/swagger.yaml) | YAML | 1,322 | 0 | 1 | 1,323 |
| [entity/biology.go](/entity/biology.go) | Go | 37 | 0 | 7 | 44 |
| [entity/company.go](/entity/company.go) | Go | 17 | 0 | 5 | 22 |
| [entity/device.go](/entity/device.go) | Go | 24 | 0 | 6 | 30 |
| [entity/fence.go](/entity/fence.go) | Go | 23 | 0 | 8 | 31 |
| [entity/role.go](/entity/role.go) | Go | 22 | 0 | 6 | 28 |
| [entity/user.go](/entity/user.go) | Go | 9 | 6 | 3 | 18 |
| [geoalgorithm/spatial_analysis_2d.go](/geoalgorithm/spatial_analysis_2d.go) | Go | 28 | 1 | 3 | 32 |
| [geocontainer/coordinate.go](/geocontainer/coordinate.go) | Go | 27 | 0 | 6 | 33 |
| [geocontainer/geo_position_2d.go](/geocontainer/geo_position_2d.go) | Go | 23 | 0 | 6 | 29 |
| [go.mod](/go.mod) | Go Module File | 78 | 0 | 5 | 83 |
| [go.sum](/go.sum) | Go Checksum File | 708 | 0 | 1 | 709 |
| [logs/2022-05-27.log](/logs/2022-05-27.log) | Log | 109 | 0 | 1 | 110 |
| [logs/2022-06-01.log](/logs/2022-06-01.log) | Log | 42 | 0 | 1 | 43 |
| [logs/2022-06-02.log](/logs/2022-06-02.log) | Log | 54 | 0 | 1 | 55 |
| [logs/2022-06-03.log](/logs/2022-06-03.log) | Log | 69 | 0 | 1 | 70 |
| [logs/2022-06-04.log](/logs/2022-06-04.log) | Log | 8 | 0 | 1 | 9 |
| [logs/2022-06-05.log](/logs/2022-06-05.log) | Log | 5 | 0 | 1 | 6 |
| [logs/2022-06-06.log](/logs/2022-06-06.log) | Log | 1 | 0 | 1 | 2 |
| [logs/2022-06-07.log](/logs/2022-06-07.log) | Log | 3 | 0 | 1 | 4 |
| [logs/2022-07-06.log](/logs/2022-07-06.log) | Log | 13 | 0 | 1 | 14 |
| [logs/2022-07-08.log](/logs/2022-07-08.log) | Log | 12 | 0 | 1 | 13 |
| [logs/2022-07-10.log](/logs/2022-07-10.log) | Log | 50 | 0 | 1 | 51 |
| [logs/2022-07-13.log](/logs/2022-07-13.log) | Log | 23 | 0 | 1 | 24 |
| [logs/2022-07-14.log](/logs/2022-07-14.log) | Log | 12 | 0 | 1 | 13 |
| [logs/2022-07-15.log](/logs/2022-07-15.log) | Log | 25 | 0 | 1 | 26 |
| [logs/2022-07-16.log](/logs/2022-07-16.log) | Log | 5 | 0 | 1 | 6 |
| [main.go](/main.go) | Go | 39 | 28 | 10 | 77 |
| [middleware/CORS_middleware.go](/middleware/CORS_middleware.go) | Go | 18 | 1 | 4 | 23 |
| [middleware/auth_middleware.go](/middleware/auth_middleware.go) | Go | 43 | 7 | 12 | 62 |
| [middleware/cache_middleware.go](/middleware/cache_middleware.go) | Go | 1 | 1 | 1 | 3 |
| [middleware/company_user_auth_middleware.go](/middleware/company_user_auth_middleware.go) | Go | 45 | 2 | 3 | 50 |
| [middleware/log_middleware.go](/middleware/log_middleware.go) | Go | 57 | 15 | 16 | 88 |
| [monitor/connection.go](/monitor/connection.go) | Go | 88 | 21 | 9 | 118 |
| [monitor/fence_jobs.go](/monitor/fence_jobs.go) | Go | 52 | 25 | 6 | 83 |
| [monitor/mesaage.go](/monitor/mesaage.go) | Go | 13 | 0 | 3 | 16 |
| [router/biology_router.go](/router/biology_router.go) | Go | 27 | 0 | 10 | 37 |
| [router/company_router.go](/router/company_router.go) | Go | 20 | 0 | 7 | 27 |
| [router/device_router.go](/router/device_router.go) | Go | 27 | 0 | 8 | 35 |
| [router/fence_router.go](/router/fence_router.go) | Go | 15 | 0 | 5 | 20 |
| [router/monitor_router.go](/router/monitor_router.go) | Go | 13 | 0 | 6 | 19 |
| [router/role_router.go](/router/role_router.go) | Go | 1 | 0 | 0 | 1 |
| [router/user_router.go](/router/user_router.go) | Go | 13 | 1 | 4 | 18 |
| [sensor/collar_message.go](/sensor/collar_message.go) | Go | 75 | 3 | 7 | 85 |
| [sensor/collar_msg_http.go](/sensor/collar_msg_http.go) | Go | 13 | 0 | 6 | 19 |
| [sensor/collections.go](/sensor/collections.go) | Go | 10 | 0 | 4 | 14 |
| [sensor/fio_message.go](/sensor/fio_message.go) | Go | 76 | 3 | 9 | 88 |
| [server/message.go](/server/message.go) | Go | 13 | 0 | 1 | 14 |
| [server/request.go](/server/request.go) | Go | 100 | 0 | 15 | 115 |
| [server/response.go](/server/response.go) | Go | 43 | 0 | 11 | 54 |
| [service/biology_service.go](/service/biology_service.go) | Go | 153 | 158 | 17 | 328 |
| [service/company_service.go](/service/company_service.go) | Go | 217 | 99 | 18 | 334 |
| [service/device_service.go](/service/device_service.go) | Go | 212 | 137 | 22 | 371 |
| [service/fence_service.go](/service/fence_service.go) | Go | 120 | 29 | 13 | 162 |
| [service/sensor_service.go](/service/sensor_service.go) | Go | 55 | 21 | 5 | 81 |
| [service/user_service.go](/service/user_service.go) | Go | 78 | 45 | 11 | 134 |
| [util/hex_to_dec.go](/util/hex_to_dec.go) | Go | 9 | 0 | 2 | 11 |
| [util/parse_date.go](/util/parse_date.go) | Go | 7 | 0 | 3 | 10 |
| [util/string_translate.go](/util/string_translate.go) | Go | 66 | 3 | 5 | 74 |
| [vo/active_fence.go](/vo/active_fence.go) | Go | 15 | 2 | 3 | 20 |
| [vo/biology_info.go](/vo/biology_info.go) | Go | 22 | 0 | 5 | 27 |
| [vo/biology_portabledevice.go](/vo/biology_portabledevice.go) | Go | 8 | 0 | 1 | 9 |
| [vo/company_info.go](/vo/company_info.go) | Go | 6 | 0 | 1 | 7 |
| [vo/company_tree_node.go](/vo/company_tree_node.go) | Go | 7 | 0 | 2 | 9 |
| [vo/fio_data.go](/vo/fio_data.go) | Go | 11 | 0 | 2 | 13 |
| [vo/fixed_device.go](/vo/fixed_device.go) | Go | 5 | 0 | 1 | 6 |
| [vo/new_collar.go](/vo/new_collar.go) | Go | 14 | 0 | 1 | 15 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)