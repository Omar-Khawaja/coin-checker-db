This is a simple Go application that calls CoinMarketCap's API to get stats 
on any cryptocurrency you choose. It then uses your Twilio account to send
you an SMS with that information. This program is a twist off the original
https://github.com/Omar-Khawaja/coin-checker and adds a database component.

This program will create the database **prices** along with a table called
**prices** if these components do not already exist. It will then store the
name of the coin, the price of the coin, and the time at which the program
was run to retrieve these statistics.

Here is an example of my querying my database after running this program a
few times with different values:

```
mysql> select * from prices;
+----------+----------+--------------------------+ 
| coin     | price    | time                     | 
+----------+----------+--------------------------+ 
| Ripple   | 0.982725 | Fri Feb 23 11:38:05 2018 | 
| Bitcoin  | 10241.0  | Fri Feb 23 11:38:13 2018 | 
| Litecoin | 206.382  | Fri Feb 23 11:38:30 2018 | 
+----------+----------+--------------------------+ 
3 rows in set (0.00 sec)
```


Please place your Twilio and db credentials in the **credentials.conf** file.

This program can be used with the *-coin* and *-info* flags.

If you do not use any flags, the price of bitcoin is returned by default.

Example:

**./main -coin=litecoin -info=all** returns the following data:

```
Name: Litecoin 
Symbol: LTC 
PriceBtc: 0.0183727 
TotalSupply: 55201008.0 
PercentChange24H: -0.24 
ID: litecoin 
Rank: 6 
Two4HVolumeUsd: 503106000.0 
MarketCapUsd: 8778174719.0 
PriceUsd: 159.022 
LastUpdated: 1518546841 
AvailableSupply: 55201008.0 
MaxSupply: 84000000.0 
PercentChange1H: 0.77 
PercentChange7D: 18.81 
```

while **./main -coin=ripple -info=price** (or simply **./main -coin=ripple**)
returns the following:

```
The price of Ripple today is $1.02957
```
