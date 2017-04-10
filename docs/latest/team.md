### Team Management Commands

The `amp team` command manages all team related operations for AMP.

```
    $ amp team

    Usage:	amp team [OPTIONS] COMMAND [ARGS...]

    Team management operations

    Options:
          --help   Print usage

    Management Commands:
      member      Team member management operations

    Commands:
      create      Create team
      get         Get team
      ls          List team
      rm          Remove team

    Run 'amp team COMMAND --help' for more information on a command.
```

### Examples

To be able to perform any team related operations, you must be logged in to AMP using a verified account.

* To create a team in an organization:
```
    $ amp team create
```

* To retrieve a list of teams:
```
    $ amp team ls
```

* To retrieve details of a specific team:
```
    $ amp team get
```

* To remove a team:
```
    $ amp team rm
    [or]
    $ amp team del
```

#### Team Member Management Commands

```
    $ amp team member

    Usage:	amp team member [OPTIONS] COMMAND [ARGS...]

    Team member management operations

    Options:
          --help   Print usage

    Commands:
      add         Add member to team
      ls          List members
      rm          Remove member from team

    Run 'amp team member COMMAND --help' for more information on a command.
```

* To add a member to a team:
```
    $ amp team member add
```

* To list members in a team:
```
    $ amp team member ls
```

* To remove a member from a team:
```
    $ amp team member rm
    [or]
    $ amp team member del
```