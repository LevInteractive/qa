# Qa

Nothing to see here. This is a **work in progress** for the moment.

## Instruction Anatomy

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

#### PRIORITY

The `PRIORITY` directive is optional and can be used to control order of
instructions within a group. This is only useful when using the same `GROUP`
within multiple files.

The lower the number, the higher priority the instructions will have. e.g. 0
will come before 1.

```
PRIORITY 3
```
