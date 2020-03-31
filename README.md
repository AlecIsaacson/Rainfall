# Rainfall and Rainfall-Detail

This pulls information from the [Northeast Ohio Regional Sewer District's rainfall dashboard](https://www.neorsd.org/stormwater-2/rainfall-dashboard/).  The dashboard is good for charting, but if you want to do more analysis, you need the underlying data.  These apps will download the data (you'll need to redirect it into a file).  You can then import it into Excel or something, and analyze it as you choose.

**Rainfall** pulls the daily totals for each day of the specified year.  It takes four command line arguments:

  -location - The city you want data for (i.e. Beachwood or "Shaker Heights")
  -year - The year you want data for.  The earliest data seems to be from 2012.
  -verbose - Optional, prints debugging info.

The list of all locations can be seen at the NEORSD rainfall dashboard.  Locations with two word names need to be quoted (see example below).

For example:

  rainfall -year=2015 -location=Beachwood
  rainfall -year=2019 -location="Shaker Heights" -verbose=true

**Rainfall-Detail** pulls the every 5 minute totals for a specific day.  It takes five arguments:

  -location - The city you want data for (i.e. Beachwood or "Shaker Heights")
  -year - The year you want data for.  The earliest data seems to be from 2012.
  -month - The month you want data for (i.e. January, Feburary, etc)
  -day - The day of the month you want data for (1, 2, 3, etc)
  -verbose - Optional, prints debugging info.

For example:

  rainfall-detail -location=Beachwood -year=2020 -month=March -day=30
