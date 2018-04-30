# Qa

1. [Installation](#installation)
2. [Getting Started](#getting-started)
3. [QA Spec](#qa-spec)
4. [Future](#future)

## Installation

```
curl -L https://github.com/LevInteractive/qa/blob/master/install.sh?raw=true | sh
```

Supported operating systems:

* OSX
* Linux
* Windows

## Getting Started

This is a simple application which allows software engineers to specify Quality
Assurance instruction along side actual code files.

**Mission**

The mission of `qa` is the following:

* Q/A spec should be versionable (e.g. checked in to git)
* Q/A spec should not take long to update and maintain
* Q/A spec should scale with your application

Take this directory structure for example:

```
├── containers
│   ├── BulkInvite.js
│   ├── .BulkInvite.qa
│   ├── CompleteRegistration.js
│   ├── .CompleteRegistration.qa
│   ├── Dashboard.js
│   ├── .Dashboard.qa
│   ├── Forgot.js
│   ├── .Forgot.qa
```

**File Generation**

Making `.qa` files hidden is optional. To generate a csv, navigate to the root
of any directory in your project and run:

```
qa . > qa.csv
```

`qa` will recurrsively walk the directory and generate a CSV based on the spec
of the files within the project.

## QA Spec

```
GROUP Owner Dashboard
DEPS User Authentication
PRIORITY 1

ACTION Select X.
EXPECT Y will happen.

ACTION Input X.
EXPECT Y will happen.

ACTION Toggle X.
EXPECT Y will happen.
```

### Definitions

These simple, yet fundamental instructions are necessary for writing clear and
efficient qa documentation.

* **Action:** Direction for executing/accomplishing something.
* **Expectation:** The outcome of an action.
* **Instructions:** One of more lines of actions and expectations.
* **Sheet:** The complete list of instructions.
* **Group:** A collection of instructions relating to similar functionality.


#### ACTION

The `ACTION` directive is used to define what action the tester should take.

```
ACTION Attempt to login with an invalid password.
```

#### EXPECT

The `EXPECT` directive is used to define what outcome an action *should* have.
An `EXPECT` should only exist after an `ACTION`.

```
EXPECT An error display explaining the password is wrong.
```

#### GROUP

The `GROUP` directive is optional and used to define the subsequent actions and
expectations within the sheet. If this directive is not provided, the
instructions will be left ungrouped (no header) at the top of the sheet.

**Multiple Files**

When using multiple files, it's valid to use the same group name. Doing so will
merge the instructions together during sheet generation.

```
GROUP User Authentication
```

#### DEPS

The `DEPS` (aka "dependencies") directive is optional and is used to provide
control over the order of `GROUP`'s. A comma-delimited list of group names which
the current group depends on may be provided.

```
DEPS User Authentication, Employee Onboarding
```

If the above is provided, it will be guarenteed that the files grouped with
"User Authentication" and "Employee Onboarding" will come first.

#### PRIORITY

The `PRIORITY` directive is optional and can be used to control order of
instructions within a group. This is only useful when using the same `GROUP`
within multiple files.

The lower the number, the higher priority the instructions will have. e.g. 0
will come before 1.

```
PRIORITY 3
```

## Future

* Google Sheets integration
* Custom dashboard integration for testers
* Listening to the community!
