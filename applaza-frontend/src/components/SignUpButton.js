import {useState} from "react";
import {Button, Form, Input, message, Modal} from "antd";
import {register} from "../utils";
import {UserOutlined} from "@ant-design/icons";

const SignUpButton = () => {
    const [modalVisible, setModalVisible] = useState(false);
    const [loading, setLoading] = useState(false);
    const handleRegisterOnClick = () => {
        setModalVisible(true);
    };
    const handleModalCancel = () => {
        setModalVisible(false);
    }
    const handleFormSubmit = async (data) => {
        setLoading(true);

        try {
            await register(data);
            message.success("Sign Up successfully");
            setModalVisible(false);
        } catch (error) {
            message.error(error.message);
        } finally {
            setLoading(false);
        }
    }
    return (
        <>
            <Button
                type="link"
                style={{padding: 0}}
                onClick={handleRegisterOnClick}
            >
                Register Now!
            </Button>
            <Modal
                title="Sign Up"
                visible={modalVisible}
                onCancel={handleModalCancel}
                footer={null}
            >
                <Form onFinish={handleFormSubmit}>
                    <Form.Item
                        name="username"
                        rules={[
                            {
                                required: true,
                                message: "Please input your username!"
                            }
                        ]}
                    >
                        <Input
                            disabled={loading}
                            prefix={<UserOutlined />}
                            placeholder="Username"
                        />
                    </Form.Item>
                    <Form.Item
                        name="password"
                        rules={[
                            {
                                required: true,
                                message: "Please input your password!"
                            }
                        ]}
                    >
                        <Input.Password
                            disabled={loading}
                            placeholder="Username"
                        />
                    </Form.Item>
                    <Form.Item>
                        <Button
                            loading={loading}
                            type="primary"
                            htmlType="submit"
                            style={{width: "100%"}}
                        >
                            Sign Up
                        </Button>
                    </Form.Item>
                </Form>
            </Modal>
        </>
    );
}

export default SignUpButton;