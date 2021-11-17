#!/bin/bash
#------------------------------------------------------
#  Get gps interface command script at MIL
#  It can be replaced at another platform if needed
#------------------------------------------------------
cell_mgmt module_info | grep 'GPS_port:' | awk '{print $2}'