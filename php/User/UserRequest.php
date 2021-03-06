<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

namespace User;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Protobuf type <code>user.UserRequest</code>
 */
class UserRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * <code>int32 id = 1;</code>
     */
    private $id = 0;
    /**
     * <code>string name = 2;</code>
     */
    private $name = '';
    /**
     * <code>string email = 3;</code>
     */
    private $email = '';

    public function __construct() {
        \GPBMetadata\User::initOnce();
        parent::__construct();
    }

    /**
     * <code>int32 id = 1;</code>
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * <code>int32 id = 1;</code>
     */
    public function setId($var)
    {
        GPBUtil::checkInt32($var);
        $this->id = $var;
    }

    /**
     * <code>string name = 2;</code>
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * <code>string name = 2;</code>
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;
    }

    /**
     * <code>string email = 3;</code>
     */
    public function getEmail()
    {
        return $this->email;
    }

    /**
     * <code>string email = 3;</code>
     */
    public function setEmail($var)
    {
        GPBUtil::checkString($var, True);
        $this->email = $var;
    }

}

