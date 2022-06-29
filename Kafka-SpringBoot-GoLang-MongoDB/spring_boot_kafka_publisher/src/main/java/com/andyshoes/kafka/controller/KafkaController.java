package com.andyshoes.kafka.controller;

import com.andyshoes.kafka.dto.Message;
import com.andyshoes.kafka.service.KafkaService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(path = "/api/v1")
public class KafkaController {

    private KafkaService kafkaService;

    @Autowired
    public KafkaController(KafkaService kafkaService){
        this.kafkaService = kafkaService;
    }

    @PostMapping
    public void insertStudent(@RequestBody Message message) {
        kafkaService.sendMessage(message);
    }
}
