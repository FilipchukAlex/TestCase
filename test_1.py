def price_convertion(price, currency):
    import requests
    from decimal import Decimal, ROUND_HALF_UP
    try:
        assert float(price) and price > 0 and currency.upper() in ('UAH', 'EUR', 'GBP')
    except (AssertionError, AttributeError, ValueError):
        print('Wrong Input Values!')
    url = f'https://free.currencyconverterapi.com/api/v5/convert?q=USD_{currency.upper()}&compact=ultra'
    try:
        requests.head(url)
    except:
        print('Something wrong!')
    exchange_rate = requests.get(url).json()
    new_price = Decimal(price) * Decimal(exchange_rate.get(f'USD_{currency.upper()}'))
    return new_price.quantize(Decimal('0.01'), rounding=ROUND_HALF_UP)


price_convertion(7.1,'gbp')
