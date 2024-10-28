# Tabeo - SpaceTrouble Booking API

Hey there! I wrote this in between errands and making dinner, so I'd expect a thorough review to bring up something.

I wrote the boilerplate in Huma just because I've been having fun using that API recently, and it takes care of documentation/validation/config for you.

There's a bunch of tradeoffs that I've made in the interest of saving time; in a full production system,
they're corners that _could_ be cut (in the interest of an MVP), but probably shouldn't be.

Examples here:

- I do not have the SpaceX api set up properly inside docker-compose.
  - It runs fine enough when you get mongodb set up for it, but the actual configuration and population of the service was taking far too long to justify it.
  - It's still configurable via `SERVICE_SPACEX_API` or `--spacex-api` in the main application, though.

- We do not have any logic around if a booking should block other bookings for the same date.
  - Logically this should happen if the booking is for a launch, but it looks a lot like we're doing space tourism here.

- A booking for one destination does not block bookings to other destinations from the same launchpad.
  - This one is more of an obvious error because it doesn't make sense (it's not physically possible), which is why I'm documenting it.

With quite a bit more time to plan I'd have properly implemented the more esoteric requirements in the readme automatically:

> Every day you change the destination for all the launchpads.
> Every day of the week from the same launchpad has to be a “flight” to a different place.

These do sound fun, and for an MVP I would have just set up destinations based on day of the week, but creating a booking already requests a flight destination from the beginning.

# Running

You can just `docker-compose up` and have a play around.
Tests still run outside of docker, and you need postgres locally to run them (they work automatically on macOS with postgres, though)