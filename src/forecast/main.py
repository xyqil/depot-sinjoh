import config
import zlib
import requests
import binascii
import csv

def ConvertKMHToMPH(primaryspeeddata):
    return primaryspeeddata * 0.621371


def ConvertMPHToKMH(secondaryspeeddata):
    return secondaryspeeddata * 1.60934


def PadAVariableWithoutZeroes(amountdata):
    return binascii.unhexlify(ZeroFillAVariable("0", amountdata))


def ConvertCelciusToFahrenheit(celcius):
    return (celcius * 1.8) + 32


def ConvertFahrenheitToCelcius(fahrenheit):
    return (fahrenheit - 32) * 0.555555555556


def ConvertTheTime(time):
    return time - 946684800 / 60


def GrabCities():
    reader = csv.dictreader(open(config.cities_csv, "r"))
    return reader


def CheckCity(cityname):
    reader = GrabCities()
    for row in reader:
        if row.city == cityname:
            return True
        elif row.city_ascii == cityname:
            return True
        else:
            return False


def GetCurrentTime():
    return now.strftime("%H:%M:%S")  # TODO: Do proper time format


def CalculateTheFileSize(buffer):
    return int(len(buffer))


class LongForecastTable:
    def __init__(self, buffer, unkdata):
        # Long Table Reference: https://github.com/RiiConnect24/Kaitai-Files/blob/master/Kaitais/forecast_file.ksy
        # Short Table Reference: https://github.com/RiiConnect24/Kaitai-Files/blob/master/Kaitais/forecast_file_short.ksy
        self.longtable = {}
        self.longtable["version"] = GetVersionByte() # TODO: Replace Version Byte
        # TODO: Figure out what file it's corresponding to, for both filesize and crc32, so as to speak and/or say.
        self.longtable["filesize"] = CalculateTheFileSize(buffer)
        self.longtable["crc32"] = zlib.crc32('buffer')
        self.longtable["opening_timestamp"] = ConvertTheTime(GetCurrentTime())
        # Unknown_1 Refrenced from https://github.com/RiiConnect24/File-Maker/blob/66d3d11e22ce3af3a6aa6a0df54b6224306e19cf/Channels/Forecast_Channel/forecast.py#L887
        self.longtable["unknown_1"] = "00" + unkdata

        self.longtable["unknown_2"] = "00" + unkdata
        self.longtable["padding"] = "00" + unkdata


def DataRequester(q, apikey, method):
    # Post code refrenced from https://www.w3schools.com/python/ref_requests_post.asp
    # API example referenced from https://www.weatherapi.com/docs/
    return requests.get(f"http://api.weatherapi.com/v1/{method}?key={apikey}&q={q}")
