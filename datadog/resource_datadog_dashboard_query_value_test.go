package datadog

import (
	"testing"
)

// JSON export used as test scenario
//{
//   "notify_list":[],
//   "description":"Created using the Datadog provider in Terraform",
//   "author_name":"--redacted--",
//   "template_variable_presets":[],
//   "template_variables":[],
//   "is_read_only":true,
//   "id":"--redacted--",
//   "title":"{{uniq}}",
//   "url":"--redacted--",
//   "created_at":"2020-10-07T21:08:56.788685+00:00",
//   "modified_at":"2020-10-07T21:08:56.788685+00:00",
//   "author_handle":"--redacted--",
//   "widgets":[
//      {
//         "definition":{
//            "custom_links":[
//               {
//                  "link":"https://app.datadoghq.com/dashboard/lists",
//                  "label":"Test Custom Link label"
//               }
//            ],
//            "autoscale":true,
//            "title":"Avg of system.mem.free over account:prod",
//            "title_align":"center",
//            "custom_unit":"Gib",
//            "precision":3,
//            "time":{
//               "live_span":"1h"
//            },
//            "title_size":"16",
//            "requests":[
//               {
//                  "q":"avg:system.mem.free{account:prod}",
//                  "aggregator":"max",
//                  "conditional_formats":[
//                     {
//                        "palette":"white_on_red",
//                        "hide_value":false,
//                        "value":9,
//                        "comparator":"<"
//                     },
//                     {
//                        "palette":"white_on_green",
//                        "hide_value":false,
//                        "value":9,
//                        "comparator":">="
//                     }
//                  ]
//               }
//            ],
//            "type":"query_value"
//         },
//         "id": "--redacted--"
//      }
//   ],
//   "layout_type":"ordered"
//}

const datadogDashboardQueryValueConfig = `
resource "datadog_dashboard" "query_value_dashboard" {
	title         = "{{uniq}}"
	description   = "Created using the Datadog provider in Terraform"
	layout_type   = "ordered"
	is_read_only  = "true"

	widget {
		query_value_definition {
			title = "Avg of system.mem.free over account:prod"
			title_align = "center"
			title_size = "16"
			custom_unit = "Gib"
			precision = "3"
			autoscale = "true"
			request {
				q = "avg:system.mem.free{account:prod}"
				aggregator = "max"
				conditional_formats {
					palette = "white_on_red"
					value = "9"
					comparator = "<"
				}
				conditional_formats {
					palette = "white_on_green"
					value = "9"
					comparator = ">="
				}
			}
			time = {
				live_span = "1h"
			}
			custom_link {
				link = "https://app.datadoghq.com/dashboard/lists"
				label = "Test Custom Link label"
			}
		}
	}
}
`

var datadogDashboardQueryValueAsserts = []string{
	"widget.0.query_value_definition.0.request.0.conditional_formats.0.comparator = <",
	"widget.0.query_value_definition.0.time.live_span = 1h",
	"widget.0.query_value_definition.0.request.0.conditional_formats.0.palette = white_on_red",
	"widget.0.query_value_definition.0.request.0.conditional_formats.0.image_url =",
	"widget.0.query_value_definition.0.precision = 3",
	"widget.0.query_value_definition.0.request.0.aggregator = max",
	"layout_type = ordered",
	"widget.0.query_value_definition.0.request.0.conditional_formats.1.palette = white_on_green",
	"widget.0.query_value_definition.0.request.0.conditional_formats.1.custom_fg_color =",
	"widget.0.query_value_definition.0.request.0.conditional_formats.1.value = 9",
	"widget.0.query_value_definition.0.autoscale = true",
	"widget.0.query_value_definition.0.request.0.q = avg:system.mem.free{account:prod}",
	"widget.0.query_value_definition.0.request.0.conditional_formats.0.custom_bg_color =",
	"widget.0.query_value_definition.0.request.0.conditional_formats.1.comparator = >=",
	"widget.0.query_value_definition.0.title_size = 16",
	"widget.0.query_value_definition.0.custom_unit = Gib",
	"widget.0.query_value_definition.0.title_align = center",
	"widget.0.query_value_definition.0.request.0.conditional_formats.0.value = 9",
	"widget.0.query_value_definition.0.request.0.conditional_formats.0.hide_value = false",
	"widget.0.query_value_definition.0.request.0.conditional_formats.0.timeframe =",
	"widget.0.query_value_definition.0.request.0.conditional_formats.1.image_url =",
	"widget.0.query_value_definition.0.request.0.conditional_formats.1.hide_value = false",
	"widget.0.query_value_definition.0.request.0.conditional_formats.1.timeframe =",
	"widget.0.query_value_definition.0.text_align =",
	"widget.0.query_value_definition.0.title = Avg of system.mem.free over account:prod",
	"widget.0.query_value_definition.0.request.0.conditional_formats.1.custom_bg_color =",
	"widget.0.query_value_definition.0.request.0.conditional_formats.# = 2",
	"description = Created using the Datadog provider in Terraform",
	"widget.0.query_value_definition.0.request.0.conditional_formats.0.custom_fg_color =",
	"is_read_only = true",
	"title = {{uniq}}",
	"widget.0.query_value_definition.0.custom_link.# = 1",
	"widget.0.query_value_definition.0.custom_link.0.label = Test Custom Link label",
	"widget.0.query_value_definition.0.custom_link.0.link = https://app.datadoghq.com/dashboard/lists",
}

func TestAccDatadogDashboardQueryValue(t *testing.T) {
	testAccDatadogDashboardWidgetUtil(t, datadogDashboardQueryValueConfig, "datadog_dashboard.query_value_dashboard", datadogDashboardQueryValueAsserts)
}

func TestAccDatadogDashboardQueryValue_import(t *testing.T) {
	testAccDatadogDashboardWidgetUtil_import(t, datadogDashboardQueryValueConfig, "datadog_dashboard.query_value_dashboard")
}
