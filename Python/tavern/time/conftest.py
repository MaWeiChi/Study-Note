#!/usr/bin/env python

import os
import urllib3

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


def pytest_tavern_beta_before_every_test_run(test_dict, variables):
    host = os.getenv("HOST", "192.168.3.8")
    port = os.getenv("PORT", "8443")

    variables["url_endpoint"] = "https://{}:{}/api/v1".format(host, port)
