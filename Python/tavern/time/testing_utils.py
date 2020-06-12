#!/usr/bin/env python

from datetime import datetime

from box import Box


def json_local_time_teardown():
    data = {
        "time": datetime.now().strftime("%Y-%m-%dT%H:%M:%S%zZ")
    }
    return Box(data)

def json_local_time_hour_teardown():
    data = {
        "time": datetime.now().strftime("%Y-%m-%dT%HZ")
    }
    return Box(data)

def json_local_date_teardown():
    data = {
        "time": datetime.now().strftime("%Y-%m-%d")
    }
    return Box(data)

def json_local_date_teardown_trans(time):
    data = {
        "time": time.split('T', 1 )[0]
    }
    return Box(data)

def compare_date_between_get_and_current(response):
    """Make sure that the response was friendly"""
    get_data=json_local_date_teardown_trans(response.json().get("data").get("time"))
    assert get_data==json_local_date_teardown()