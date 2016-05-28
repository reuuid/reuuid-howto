{-# LANGUAGE OverloadedStrings #-}

import Control.Monad        (void)

import Data.Char            (ord)
import Data.ByteString.Lazy (ByteString, intercalate, split)

import Network.HTTP.Client


getUUID :: IO ByteString
getUUID = do
  manager <- newManager defaultManagerSettings
  request <- parseUrl "http://reuuid.org/get/"
  response <- httpLbs request manager
  return $ responseBody response

getUUIDs :: Int -> IO [ByteString]
getUUIDs n
  | n < 1 = return []
  | otherwise = do
      manager <- newManager defaultManagerSettings
      request <- parseUrl ("http://reuuid.org/get/" ++ show n)
      response <- httpLbs request manager
      return $ init $ split (fromIntegral (ord '\n')) $ responseBody response

donateUUIDs :: [ByteString] -> IO ()
donateUUIDs [] = return ()
donateUUIDs uuids = do
  manager <- newManager defaultManagerSettings
  initialRequest <- parseUrl "http://reuuid.org/give/"
  let body = intercalate "\n" uuids
  let request = initialRequest { method = "POST"
                               , requestBody = RequestBodyLBS body
                               }
  void $ httpLbs request manager


main :: IO ()
main = do
  one <- getUUID
  five <- getUUIDs 5
  mapM_ print (one : five)
  donateUUIDs (["3b54969c-d9fa-4ac9-aa38-4c69590ebaa5"
               ,"1237168c-35c4-437f-94fe-f48fe972eafa"
               ,"ac1c7d22-4903-4231-ae9a-c042c3a6211d"
               ,"b63f76ac-3d7c-43fd-b966-38ce938a126e"
               ,"03e65fe3-21d2-4ef4-bbf1-f14bb42f06e3"
               -- give back the ones we took (don't be greedy)
               ] ++ one : five)
