<?php
/**
 * Created by PhpStorm.
 * User: puzanov
 * Date: 7/7/17
 * Time: 7:55 PM
 */

$url = "https://www.instagram.com/hippohae/?__a=1";
$json = file_get_contents($url);
$data = json_decode($json);
var_dump($data);
var_dump($data->user->followed_by);