# iOffice CLI Reservation Tool

Lists and creates reservations of spaces (desks/rooms) in iOffice

_This tool is not created by, affiliate with, or supported by iOffice_

![screenshot.png](screenshot.png)


## Usage

This section assumes you have downloaded the application from the latest [release](https://github.com/alicekaerast/ioffice/releases/latest) and have installed this to somewhere in $PATH. If you are running this from source code, replace `ioffice` with `go run .` in the below documentation.

1. create ioffice.yaml in either your $HOME directory or ~/.config/ (or this directory if working from source code)
2. Set up your authentication (see Auth section next)
3. Set your instance's hostname (`example` in https://example.ioffice.com)
4. Run the application to get a list of your future bookings (`ioffice`)
5. Book a named room by passing in a date and a room name (`ioffice create 2022-03-13 2101`)
6. From your list of reservations, take the room ID of the desk you prefer and put it into ioffice.yaml
7. Now you can also reserve this desk for a full day without passing in the number (`ioffice . create 2022-03-13`)
8. You can then cancel or check in to your reservation by passing its ID (`ioffice checkin 68610` or `ioffice cancel 68610`)
9. You can also get a list of floors for a building (`ioffice floors [buildingID]`) which will be useful for some planned future features

### Auth

If you do not use SSO then you just need to add your username and password to ioffice.yaml, otherwise read on.

IOffice API currently can only handle simple auth.  However, the frontend can actually use an SSO provider.
If you have a non-SSO account then change `username` and `password` to your username and password.
Otherwise you will need to manually login via the UI to your SSO provider and pass the authentication.
Once this has been done, inspect your cookies for `ACTID` - set this as your `session`.  This will
allow this application to work even with SSO providing the session is active.

*Note* We prioritize `session` over `username` and `password`.  So if you have a non-SSO account and want to avoid 
`session` then ensure it is set to blank `""` value.

### Duplicate Room Names

You might have the same room names in multiple offices. If this is the case then you can set the building that you work in.

1. Run `ioffice buildings` to get a list of building IDs and names
2. Set `buildingID` in the yaml file to the ID of the building you work in

All reservations will now use only this building

## Development

iOffice publish [API documentation](https://ioffice.github.io/api/) that helps figure out new functionality. It's inaccurate or plain wrong in places, but you can use dev tools in your browser to watch what API calls the web app makes. If you feel like fixing the documentation, you can [raise a PR](https://github.com/iOffice/ioffice.github.io) against it.

To build and install locally: `make install` (assuming you have Go already installed). You may need to do this if your Mac blocks unsigned applications from running.

To release a new version:

1. Get your code into the main branch (via a PR if you're not a collaborator, ideally conventional commit messages)
2. Create a git tag with the new desired version number (follow SemVer)
3. Push your git tag (`git push --tags`)
4. Go Releaser will create a new release
5. If you're happy the release is stable, change it from a pre-release to a full release
