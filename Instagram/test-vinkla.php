<?php
/**
 * Created by PhpStorm.
 * User: puzanov
 * Date: 7/7/17
 * Time: 7:40 PM
 */

require_once 'vendor/autoload.php';
use Vinkla\Instagram\Instagram;
$instagram = new Instagram();
$instagram->get('hippohae');
