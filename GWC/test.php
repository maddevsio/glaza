<?php

require_once './vendor/autoload.php';

try{

    $client = new Google_Client();
    $client->setApplicationName('GWC');
    $client->setAuthConfig('./organic-search-analytics-2ee3b6f08d9d.json');
    $client->setRedirectUri('urn:ietf:wg:oauth:2.0:oob');
    $client->setScopes(['https://www.googleapis.com/auth/webmasters.readonly']);
    $service = new Google_Service_Webmasters($client);
    $t = $service->sites->get('https://silkroadexplore.com/');
    print_r($t);

    $t = $service->searchanalytics->query('https://silkroadexplore.com/');
    print_r($t);


}
catch(Google_Exception $e){
    echo $e->getMessage();
}
catch(Google_Service_Exception $e){
    echo $e->getMessage();
}
