<?php
/**
 * Created by PhpStorm.
 * User: puzanov
 * Date: 7/7/17
 * Time: 7:25 PM
 */

require_once 'vendor/autoload.php';
$res = Bolandish\Instagram::getMediaByUserID("hippohae");
var_dump($res);