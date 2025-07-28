package saasruntime_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
)

func TestAccSaasRuntimeUnitKind_saasRuntimeUnitKindBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckSaasRuntimeUnitKindDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccSaasRuntimeUnitKind_saasRuntimeUnitKindBasicExample_basic(context),
			},
			{
				ResourceName:            "google_saas_runtime_unit_kind.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "labels", "location", "terraform_labels", "unit_kind_id"},
			},
			{
				Config: testAccSaasRuntimeUnitKind_saasRuntimeUnitKindBasicExample_update(context),
			},
			{
				ResourceName:            "google_saas_runtime_unit_kind.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "labels", "location", "terraform_labels", "unit_kind_id"},
			},
		},
	})
}

func testAccSaasRuntimeUnitKind_saasRuntimeUnitKindBasicExample_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_saas_runtime_saas" "test_saas" {
  provider = google-beta
  saas_id  = "tf-test-test-saas%{random_suffix}"
  location = "global"

  locations {
    name = "us-central1"
  }
  locations {
    name = "europe-west1"
  }
}

resource "google_saas_runtime_unit_kind" "dependency" {
  provider        = google-beta
  unit_kind_id    = "tf-test-test-unit-kind%{random_suffix}-dependency"
  location        = "global"
  saas            = google_saas_runtime_saas.test_saas.id
}

resource "google_saas_runtime_unit_kind" "example" {
  provider        = google-beta
  unit_kind_id    = "tf-test-test-unit-kind%{random_suffix}"
  location        = "global"
  saas            = google_saas_runtime_saas.test_saas.id
  dependencies {
    alias     = "dep1"
    unit_kind = google_saas_runtime_unit_kind.dependency.id
  }
}
`, context)
}

func testAccSaasRuntimeUnitKind_saasRuntimeUnitKindBasicExample_update(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_saas_runtime_saas" "test_saas" {
  provider = google-beta
  saas_id  = "tf-test-test-saas%{random_suffix}"
  location = "global"

  locations {
    name = "us-central1"
  }
  locations {
    name = "europe-west1"
  }
}

resource "google_saas_runtime_unit_kind" "dependency" {
  provider        = google-beta
  unit_kind_id    = "tf-test-test-unit-kind%{random_suffix}-dependency"
  location        = "global"
  saas            = google_saas_runtime_saas.test_saas.id
}

resource "google_saas_runtime_unit_kind" "example" {
  provider        = google-beta
  unit_kind_id    = "tf-test-test-unit-kind%{random_suffix}"
  location        = "global"
  saas            = google_saas_runtime_saas.test_saas.id
  dependencies {
    alias     = "dep1"
    unit_kind = google_saas_runtime_unit_kind.dependency.id
  }
  input_variable_mappings {
    variable = "input_var"
    to {
      dependency     = "dep1"
      input_variable = "dep_input_var"
    }
  }
  output_variable_mappings {
    variable = "output_var"
    from {
      dependency      = "dep1"
      output_variable = "dep_output_var"
    }
  }
}
`, context)
}
