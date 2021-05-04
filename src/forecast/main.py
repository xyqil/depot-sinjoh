import config
import zlib
import requests
import binascii
import collections
import csv

class AbstractionLayer:
  
	def CalcCRC32(filename):
	    # From: https://stackoverflow.com/questions/1742866/compute-crc-of-file-in-python
	    with open(filename, 'rb') as fh:
		hash = 0
		while True:
		    s = fh.read(65536)
		    if not s:
			break
		    hash = zlib.crc32(s, hash)
		return "%08X" % (hash & 0xFFFFFFFF)

	def GetVersionByte():
		return binascii.unhexlify('30')

	def ZeroFillAVariable(variabledata, amount):  
		# Thanks Snoot for figuring out a better name for the function
		return variabledata.zfill(amount)
    
	def DataDuplicator(unduplicateddata, integerdata):
		return unduplicateddata * integerdata

	def ConvertKMHToMPH(primaryspeeddata):
		return primaryspeeddata * 0.621371

	def ConvertMPHToKMH(secondaryspeeddata):
		return secondaryspeeddata * 1.60934
  
	def PadAVariableWithoutZeroes(amountdata):
		return binascii.unhexlify(ZeroFillAVariable('0', 2))

  	def ConvertCelciusToFahrenheit(celcius):
    		return (celcius * 1.8) + 32
  
  	def ConvertFahrenheitToCelcius(fahrenheit):
    		return (fahrenheit - 32) * 0.555555555556  
	
  	def time_convert(time):
    		# This function and nothing else was made by Larsenv the Seagull, modified so it isn't pure crap :troll:
    		return time - 946684800 / 60
  
  	def GrabCities():
  		reader = csv.dictreader(open(config.cities_csv, "r"))
  		return reader

  	def CheckCity(cityname):
    	reader = GrabCities()
    	for row in reader:
      		if row."city" == cityname:
        		return True
      		if row."city_ascii" == cityname:
        		return True
		else:
    			return False

class Tables:
	def Bellossom():
		longtable = Collections.OrderedDict()
