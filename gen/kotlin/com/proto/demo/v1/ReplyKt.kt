// Generated by the protocol buffer compiler. DO NOT EDIT!
// source: proto/demo/v1/demo.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package com.proto.demo.v1;

@kotlin.jvm.JvmName("-initializereply")
public inline fun reply(block: com.proto.demo.v1.ReplyKt.Dsl.() -> kotlin.Unit): com.proto.demo.v1.Reply =
  com.proto.demo.v1.ReplyKt.Dsl._create(com.proto.demo.v1.Reply.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `proto.demo.v1.Reply`
 */
public object ReplyKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: com.proto.demo.v1.Reply.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: com.proto.demo.v1.Reply.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): com.proto.demo.v1.Reply = _builder.build()

    /**
     * `.proto.demo.v1.IMPORTANCE priority = 1 [json_name = "priority"];`
     */
    public var priority: com.proto.demo.v1.IMPORTANCE
      @JvmName("getPriority")
      get() = _builder.getPriority()
      @JvmName("setPriority")
      set(value) {
        _builder.setPriority(value)
      }
    public var priorityValue: kotlin.Int
      @JvmName("getPriorityValue")
      get() = _builder.getPriorityValue()
      @JvmName("setPriorityValue")
      set(value) {
        _builder.setPriorityValue(value)
      }
    /**
     * `.proto.demo.v1.IMPORTANCE priority = 1 [json_name = "priority"];`
     */
    public fun clearPriority() {
      _builder.clearPriority()
    }

    /**
     * `string from = 2 [json_name = "from"];`
     */
    public var from: kotlin.String
      @JvmName("getFrom")
      get() = _builder.getFrom()
      @JvmName("setFrom")
      set(value) {
        _builder.setFrom(value)
      }
    /**
     * `string from = 2 [json_name = "from"];`
     */
    public fun clearFrom() {
      _builder.clearFrom()
    }

    /**
     * `string to = 3 [json_name = "to"];`
     */
    public var to: kotlin.String
      @JvmName("getTo")
      get() = _builder.getTo()
      @JvmName("setTo")
      set(value) {
        _builder.setTo(value)
      }
    /**
     * `string to = 3 [json_name = "to"];`
     */
    public fun clearTo() {
      _builder.clearTo()
    }

    /**
     * `string subject = 4 [json_name = "subject"];`
     */
    public var subject: kotlin.String
      @JvmName("getSubject")
      get() = _builder.getSubject()
      @JvmName("setSubject")
      set(value) {
        _builder.setSubject(value)
      }
    /**
     * `string subject = 4 [json_name = "subject"];`
     */
    public fun clearSubject() {
      _builder.clearSubject()
    }

    /**
     * `string body = 5 [json_name = "body"];`
     */
    public var body: kotlin.String
      @JvmName("getBody")
      get() = _builder.getBody()
      @JvmName("setBody")
      set(value) {
        _builder.setBody(value)
      }
    /**
     * `string body = 5 [json_name = "body"];`
     */
    public fun clearBody() {
      _builder.clearBody()
    }
  }
}
@kotlin.jvm.JvmSynthetic
@com.google.errorprone.annotations.CheckReturnValue
public inline fun com.proto.demo.v1.Reply.copy(block: com.proto.demo.v1.ReplyKt.Dsl.() -> kotlin.Unit): com.proto.demo.v1.Reply =
  com.proto.demo.v1.ReplyKt.Dsl._create(this.toBuilder()).apply { block() }._build()

