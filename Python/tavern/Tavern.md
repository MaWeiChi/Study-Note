# Tavern

## Strict key checking:

### Example

The response from a server is `[ 1, 2, 3 ]` :

**strict: True (default)**

pass:

- `[1, 2, 3]`

**strict: False**

pass:

- `[1, 2, 3], [1], [2], [3], [1,2], [1,3], [2,3]`

**both not pass:**

- `[3, 1]`, `[2, 1]` - items present, but out of order
- `[2, 4]` - ‘4’ not present in response from the server

## **Changing the setting**

This setting can be controlled in 3 different ways, the order of priority being:

1. In the test/stage itself
2. Passed on the command line
3. Read from pytest config

This means that using the command line option will *not* override any settings for specific tests.

Each of these methods is done by passing a sequence of strings indicating which section (`json`/`redirect_query_params`/`headers`) should be affected, and optionally whether it is on or off.

- `json:off headers:on` - turn off for the body, but on for the headers. `redirect_query_params` will stay default off.
- `json:off headers:off` - turn body and header strict checking off
- `redirect_query_params:on json:on` redirect parameters is turned on and json is kept on (as it is on by default), header strict matching is kept off (as default).

Leaving the ‘on’ or ‘off’ at the end of each setting will imply ‘on’ - ie, using `json headers redirect_query_params` as an option will turn them all on.

### **Command line**

There is a command line argument, `--tavern-strict`, which controls the default global strictness setting.

```bash
# Enable strict checking for body and headers only
py.test --tavern-strict json:on headers:on redirect_query_params:off -- my_test_folder/
```

### **In the Pytest config file**

This behaves identically to the command line option, but will be read from whichever configuration file Pytest is using.

```python
[pytest]
tavern-strict=json:off headers:on
```

### **Per test**

Strictness can also be enabled or disabled on a per-test basis. The `strict` key at the top level of the test should a list consisting of one or more strictness setting as described in the previous section.

```yaml
---

test_name: Make sure the headers match what I expect exactly

strict:
  - headers:on
  - json:off

stages:
  - name: Try to get user
    request:
      url: "{host}/users/joebloggs"
      method: GET
    response:
      status_code: 200
      headers:
        content-type: application/json
        content-length: 20
        x-my-custom-header: chocolate
      json:
        # As long as "id: 1" is in the response, 
        # this will pass and other keys will be ignored
        id: 1
```

A special option that can be done at the test level (or at the stage level, as described in the next section) is just to pass a boolean. This will turn strict checking on or off for all settings for the duration of that test/stage.

```yaml
test_name: Just check for one thing in a big nested dict

# completely disable strict key checking for this whole test
strict: False

stages:
  - name: Try to get user
    request:
      url: "{host}/users/joebloggs"
      method: GET
    response:
      status_code: 200
      json:
        q:
          x:
            z:
              a: 1
```

### **Per stage**

Often you have a standard stage before other stages, such as logging in to your server, where you only care if it returns a 200 to indicate that you’re logged in. To facilitate this, you can enable or disable strict key checking on a per-stage basis as well.

Two examples for doing this - these examples should behave identically:

```yaml
---

# Enable strict checking for this test, but disable it for the login stage

test_name: Login and create a new user

# Force re-enable strict checking, in case it was turned off globally
strict:
  - json:on

stages:
  - name: log in
    request:
      url: "{host}/users/joebloggs"
      method: GET
    response:
      # Disable all strict key checking just for this stage
      strict: False
      status_code: 200
      json:
        logged_in: True
        # Ignores any extra metadata like user id, last login, etc.

  - name: Create a new user
    request:
      url: "{host}/users/joebloggs"
      method: POST
      json: &create_user
        first_name: joe
        last_name: bloggs
        email: joe@bloggs.com
    response:
      status_code: 200
      # Because strict was set 'on' at the test level, this must match exactly
      json:
        <<: *create_user
        id: 1
```

Or if strict json key checking was enabled at the global level:

```yaml
---

test_name: Login and create a new user

stages:
  - name: log in
    request:
      url: "{host}/users/joebloggs"
      method: GET
    response:
      strict:
        - json:off
      status_code: 200
      json:
        logged_in: True

  - name: Create a new user
    request: ...
```



## Share data:

### file > test > stage > name

**In same file:**

Define an anchor using `&name_of_anchor`.

This can then be assigned to another object using `new_object: *name_or_anchor`, or they can be used to extend objects using `<<: *name_of_anchor`.

```yaml
# input.yaml
---
first: &top_anchor
  a: b
  c: d

second: *top_anchor

third:
  <<: *top_anchor
  c: overwritten
  e: f
```

If we convert this to JSON, for example with a script like this:

```python
#!/usr/bin/env python

# load.py
import yaml
import json

with open("input.yaml", "r") as yfile:
    for doc in yaml.load_all(yfile.read()):
        print(json.dumps(doc, indent=2))
```

Output:

```python
{
  'first': {
    'a': 'b',
    'c': 'd'
  },
  'second': {
    'a': 'b',
    'c': 'd'
  },
  'third': {
    'a': 'b',
    'c': 'overwritten',
    'e': 'f'
  }
}
```