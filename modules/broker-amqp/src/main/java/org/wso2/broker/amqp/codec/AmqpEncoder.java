/*
 * Copyright (c) 2017, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package org.wso2.broker.amqp.codec;

import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandlerContext;
import io.netty.handler.codec.MessageToMessageEncoder;
import org.wso2.broker.amqp.codec.frames.GeneralFrame;

import java.util.List;

/**
 * Netty based AMQP encoder.
 */
public class AmqpEncoder extends MessageToMessageEncoder<GeneralFrame> {

    protected void encode(ChannelHandlerContext ctx, GeneralFrame frame, List<Object> out)
            throws Exception {
        ByteBuf buf = frame.write(ctx.alloc());
        out.add(buf);
    }
}
