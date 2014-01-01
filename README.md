Fairsplit
=========

Fairsplit is a program that helps you to split expenses fairly among a group
of people.

It takes a list of transactions from the standard input until EOF is
encountered (which can be inserted with %[2]s on this operating system.)

Transactions are entered by showing how money has flowed from one person on
behalf of others. For example, the transaction

    Ben 40.00 Ben Lila Carlos Emil

means that Ben paid 40.00 of whatever currency on behalf of Ben, Lila, Carlos,
and Emil, so that Lila, Carlos, and Emil each owe Ben 10.00.

Fairsplit also minimizes the number of transactions, for example given the
following input, only Sonia pays Reed anything:

    > Ben 45.67 Ben Reed Sonia
    > Reed 78 Sonia Ben
    > Sonia 33.2 Sonia Ben Reed
    > Ben 19.62 Sonia
    > <%[2]s>

    OUTSTANDING TRANSACTIONS:
    Sonia must pay:
        51.71 to Reed

    THANK YOU.

Installation
------------

You can download binaries for the current release, they can be run without
installation.

You can also download the source code and compile it yourself. See the
[Go](golang.org) website for more information on the Go programming
language, and how to get and compile the code here.
