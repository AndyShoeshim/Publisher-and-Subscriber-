package com.andyshoes.kafka.service;

import com.andyshoes.kafka.dto.Message;
import org.springframework.context.annotation.Bean;
import org.springframework.kafka.core.KafkaTemplate;

public class KafkaService {

    private KafkaTemplate<String, Message> kafkaTemplate;

    public KafkaService(KafkaTemplate<String, Message> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    public void sendMessage(Message message) {
        kafkaTemplate.send("myTopic", message);
    }
}
