<?php
/**
 * Created by PhpStorm.
 * User: puzanov
 * Date: 7/7/17
 * Time: 7:55 PM
 */

require_once __DIR__ . "/vendor/autoload.php";

use Monolog\Logger;
use Monolog\Handler\StreamHandler;

$logger = new \Monolog\Logger("log");
$logger->pushHandler(new StreamHandler('php://stdout', Logger::DEBUG));

$client = new \MongoDB\Client("mongodb://mongodb:27017");
$collection = $client->glaza->instagram;

while (true) {
  $accounts = array();
  if (getenv('INSTAGRAM_ACCOUNTS')) {
    $accounts = explode(',', getenv('INSTAGRAM_ACCOUNTS'));
  } else {
    $logger->addError("No Instagram accounts");
    exit;
  }

  foreach ($accounts as $account) {
    $url = "https://www.instagram.com/$account/?__a=1";

    $json = file_get_contents($url);
    $data = json_decode($json);
    if (json_last_error() != JSON_ERROR_NONE) {
      $logger->addError("JSON is not valid. Skip saving...");
      continue;
    }
    $result = $collection->insertOne($data);
    $logger->addInfo("mongo result '{$result->getInsertedId()}");
    $logger->addInfo("Instagram followed_by " . $data->user->followed_by->count);
    $logger->addInfo("Instagram follows " . $data->user->follows->count);
    $logger->addInfo("Instagram photos " . $data->user->media->count);
  }
  
  sleep(60);
}