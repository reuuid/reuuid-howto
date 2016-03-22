<?php
    // To Get One
    $uuid = file_get_contents('http://reuuid.org/get/');
    print($uuid . "\n");

    // To Get Five
    $uuids = file_get_contents('http://reuuid.org/get/5');
    print($uuids . "\n");

    // To Give Some
    function donate($uuids) {
        // Setup a post request
        $body = implode($uuids, "\n");
        $httpReq = array('http' => array('method' => 'POST',
                                        'content' => $body));

        // Create a stream context and send it
        $ctx = stream_context_create($httpReq);
        $f = @fopen('http://reuuid.org/give/', 'r', false, $ctx);
        if (!$f)
            throw new Exception('Couldn\'t send the Uuids.');
    }
    
    donate(array(
        '3b54969c-d9fa-4ac9-aa38-4c69590ebaa5',
        '1237168c-35c4-437f-94fe-f48fe972eafa',
        'ac1c7d22-4903-4231-ae9a-c042c3a6211d',
        'b63f76ac-3d7c-43fd-b966-38ce938a126e',
        '03e65fe3-21d2-4ef4-bbf1-f14bb42f06e3',
    ));
?>

