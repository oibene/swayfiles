export default function ModalAccount(){
    return(
        <div className="absolute bg-dark-gray w-80 h-50">
            <div className="font-noto font-bold text-light-gray text-sm">

                <div className="flex m-5">
                    <div className="h-15 w-15 rounded-full bg-light-gray"></div>
                    <div className="ml-4">
                        <p>CONTA</p>
                    </div>

                </div>
                <hr className="w-full h-0.5 rounded-sm"/>
                <p>ALTERAR CONTA</p>
                <p>LOGOUT</p>

            </div>

        </div>
    )
}