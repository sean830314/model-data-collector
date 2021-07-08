import logging
import datetime
import json
from pythonjsonlogger import jsonlogger


class CustomJsonFormatter(jsonlogger.JsonFormatter):
    def add_fields(self, log_record, record, message_dict):
        super(CustomJsonFormatter, self).add_fields(log_record, record, message_dict)
        if not log_record.get('timestamp'):
            # this doesn't use record.created, so it is slightly off
            now = datetime.datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%S.%fZ')
            log_record['timestamp'] = now
        if log_record.get('level'):
            log_record['level'] = log_record['level'].upper()
        else:
            log_record['level'] = record.levelname
            
logger = logging.getLogger()
logger.setLevel(logging.DEBUG)
formatter = CustomJsonFormatter('%(timestamp)s %(level)s %(name)s %(message)s')

stream_logHandler = logging.StreamHandler()
stream_logHandler.setLevel(logging.DEBUG)
stream_logHandler.setFormatter(formatter)

log_filename = datetime.datetime.now().strftime("%Y-%m-%d_%H_%M_%S.log")
file_logHandler = logging.FileHandler("data/{}".format(log_filename))
file_logHandler.setLevel(logging.DEBUG)
file_logHandler.setFormatter(formatter)
logger.addHandler(stream_logHandler)
logger.addHandler(file_logHandler)
import time
st = time.time()
while time.time() - st < 20000:
    payload = {"user_id": "60e55e32f321e00001de82c5", "version_id": "60e56863b751572cbd49e0b1", "data": [{"paragraph": "Paperupdates  https://docs.google.com/document/d/1Z4M_qwILLRehPbVRUsJ3OF8Iir-gqS-ZYe7W-LE9gnE/edit#heading=h.m6iml6hqrnm2  Hyperledger  Whitepaper  Abstract  This paper describes industry use cases that drive the principles behind a  newblockchain fabric, and outlines the basic requirements and highlevel architecture  basedon those use cases. The design presented here describes this evolving blockchainfabric, calledHyperledger,  as a protocol for businesstobusiness andbusinesstocustomer transactions. Hyperledger  allows for compliance with regulations,while supporting the varied requirements that arise when competing  businesses worktogether on the same network. The central elements of this specification  (describedbelow) are smart contracts (a.k.a. chaincode), digital assets, record  repositories, adecentralized consensusbased network, and cryptographic security. To  theseblockchain staples, industry requirements for performance, verified identities,  privateand confidential transactions, and a pluggable consensus model have been  added.", "targets_id": ["5", "21"]}]}
    logger.info("model_predicts: {}".format(json.dumps(payload)))
    time.sleep(10)
    print("sss")